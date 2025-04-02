#import "notifications_darwin.h"
#include <Foundation/Foundation.h>

#if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
#import <UserNotifications/UserNotifications.h>
#endif

bool isNotificationAvailable(void) {
    if (@available(macOS 11.0, *)) {
        return true; // UserNotifications available
    } else if (@available(macOS 10.8, *)) {
        NSLog(@"UNUserNotification API requires macOS 11.0 or later, falling back to NSUserNotification.");
        return true; // NSUserNotification available
    }
    return false; // Neither API available
}

bool checkBundleIdentifier(void) {
    NSBundle *main = [NSBundle mainBundle];
    if (main.bundleIdentifier == nil) {
        return NO;
    }
    return YES;
}

extern void captureResult(int channelID, bool success, const char* error);
extern void didReceiveNotificationResponse(const char *jsonPayload, const char* error);

#if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
@interface NotificationsDelegate : NSObject <UNUserNotificationCenterDelegate>
@end

@implementation NotificationsDelegate

- (void)userNotificationCenter:(UNUserNotificationCenter *)center
       willPresentNotification:(UNNotification *)notification
         withCompletionHandler:(void (^)(UNNotificationPresentationOptions options))completionHandler {
    UNNotificationPresentationOptions options = 0;
    
    if (@available(macOS 11.0, *)) {
        // These options are only available in macOS 11.0+
        options = UNNotificationPresentationOptionList | 
                  UNNotificationPresentationOptionBanner | 
                  UNNotificationPresentationOptionSound;
    }
    
    completionHandler(options);
}

- (void)userNotificationCenter:(UNUserNotificationCenter *)center
didReceiveNotificationResponse:(UNNotificationResponse *)response
         withCompletionHandler:(void (^)(void))completionHandler {

    NSMutableDictionary *payload = [NSMutableDictionary dictionary];
    
    [payload setObject:response.notification.request.identifier forKey:@"id"];
    [payload setObject:response.actionIdentifier forKey:@"actionIdentifier"];
    [payload setObject:response.notification.request.content.title ?: @"" forKey:@"title"];
    [payload setObject:response.notification.request.content.body ?: @"" forKey:@"body"];
    
    if (response.notification.request.content.categoryIdentifier) {
        [payload setObject:response.notification.request.content.categoryIdentifier forKey:@"categoryIdentifier"];
    }

    if (response.notification.request.content.subtitle) {
        [payload setObject:response.notification.request.content.subtitle forKey:@"subtitle"];
    }
    
    if (response.notification.request.content.userInfo) {
        [payload setObject:response.notification.request.content.userInfo forKey:@"userInfo"];
    }
    
    if ([response isKindOfClass:[UNTextInputNotificationResponse class]]) {
        UNTextInputNotificationResponse *textResponse = (UNTextInputNotificationResponse *)response;
        [payload setObject:textResponse.userText forKey:@"userText"];
    }
    
    NSError *error = nil;
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload options:0 error:&error];
    if (error) {
        NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [error localizedDescription]];
        didReceiveNotificationResponse(NULL, [errorMsg UTF8String]);
    } else {
        NSString *jsonString = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
        didReceiveNotificationResponse([jsonString UTF8String], NULL);
    }
    
    completionHandler();
}

@end

static NotificationsDelegate *delegateInstance = nil;
static dispatch_once_t onceToken;

static BOOL ensureDelegateInitialized(void) {
    __block BOOL success = YES;

    dispatch_once(&onceToken, ^{
        delegateInstance = [[NotificationsDelegate alloc] init];
        UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
        center.delegate = delegateInstance;
    });

    if (!delegateInstance) {
        success = NO;
    }

    return success;
}
#endif


#if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080

static NSString *const kCategoriesKey = @"NotificationCategories";

@interface LegacyNotificationsDelegate : NSObject <NSUserNotificationCenterDelegate>
@property (nonatomic, strong) NSMutableDictionary *pendingNotifications;
@end

@implementation LegacyNotificationsDelegate

- (instancetype)init {
    self = [super init];
    if (self) {
        self.pendingNotifications = [NSMutableDictionary dictionary];
        [[NSUserNotificationCenter defaultUserNotificationCenter] setDelegate:self];
    }
    return self;
}


- (BOOL)userNotificationCenter:(NSUserNotificationCenter *)center shouldPresentNotification:(NSUserNotification *)notification {
    return YES; // Always show notifications
}

- (void)userNotificationCenter:(NSUserNotificationCenter *)center didActivateNotification:(NSUserNotification *)notification {
    NSMutableDictionary *payload = [NSMutableDictionary dictionary];
    
    [payload setObject:notification.identifier ?: @"" forKey:@"id"];
    [payload setObject:notification.title ?: @"" forKey:@"title"];
    [payload setObject:notification.subtitle ?: @"" forKey:@"subtitle"];
    [payload setObject:notification.informativeText ?: @"" forKey:@"body"];
    
    // Handle action
    NSString *actionIdentifier = @"DEFAULT_ACTION";
    if (notification.activationType == NSUserNotificationActivationTypeActionButtonClicked) {
        actionIdentifier = @"ACTION_BUTTON";
    } else if (notification.activationType == NSUserNotificationActivationTypeReplied) {
        actionIdentifier = @"TEXT_REPLY";
        if (notification.response.string) {
            [payload setObject:notification.response.string forKey:@"userText"];
        }
    }
    [payload setObject:actionIdentifier forKey:@"actionIdentifier"];
    
    // Include user info if present
    if (notification.userInfo) {
        [payload setObject:notification.userInfo forKey:@"userInfo"];
    }
    
    NSError *error = nil;
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload options:0 error:&error];
    if (error) {
        NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [error localizedDescription]];
        didReceiveNotificationResponse(NULL, [errorMsg UTF8String]);
    } else {
        NSString *jsonString = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
        didReceiveNotificationResponse([jsonString UTF8String], NULL);
    }
}

@end

static LegacyNotificationsDelegate *legacyDelegateInstance = nil;
static dispatch_once_t legacyOnceToken;

static BOOL ensureLegacyDelegateInitialized(void) {
    __block BOOL success = YES;

    dispatch_once(&legacyOnceToken, ^{
        legacyDelegateInstance = [[LegacyNotificationsDelegate alloc] init];
    });

    if (!legacyDelegateInstance) {
        success = NO;
    }

    return success;
}

// Function to save a notification category
void saveNotificationCategory(NSString *categoryId, NSDictionary *categoryData) {
    NSUserDefaults *defaults = [NSUserDefaults standardUserDefaults];
    
    // Get existing categories dictionary
    NSMutableDictionary *categories = [NSMutableDictionary dictionaryWithDictionary:
                                      [defaults objectForKey:kCategoriesKey] ?: @{}];
    
    // Add or update the category
    categories[categoryId] = categoryData;
    
    // Save back to user defaults
    [defaults setObject:categories forKey:kCategoriesKey];
    [defaults synchronize];
}

// Function to retrieve a notification category
NSDictionary* getNotificationCategory(NSString *categoryId) {
    NSUserDefaults *defaults = [NSUserDefaults standardUserDefaults];
    NSDictionary *categories = [defaults objectForKey:kCategoriesKey];
    return categories[categoryId];
}

// Function to remove a notification category
void removeStoredNotificationCategory(NSString *categoryId) {
    NSUserDefaults *defaults = [NSUserDefaults standardUserDefaults];
    
    // Get existing categories dictionary
    NSMutableDictionary *categories = [NSMutableDictionary dictionaryWithDictionary:
                                      [defaults objectForKey:kCategoriesKey] ?: @{}];
    
    // Remove the category
    [categories removeObjectForKey:categoryId];
    
    // Save back to user defaults
    [defaults setObject:categories forKey:kCategoriesKey];
    [defaults synchronize];
}
#endif


void requestNotificationAuthorization(int channelID) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            if (!ensureDelegateInitialized()) {
                NSString *errorMsg = @"Notification delegate has been lost. Reinitialize the notification service.";
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            UNAuthorizationOptions options = UNAuthorizationOptionAlert | UNAuthorizationOptionSound | UNAuthorizationOptionBadge;
            
            [center requestAuthorizationWithOptions:options completionHandler:^(BOOL granted, NSError * _Nullable error) {
                if (error) {
                    NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [error localizedDescription]];
                    captureResult(channelID, false, [errorMsg UTF8String]);
                } else {
                    captureResult(channelID, granted, NULL);
                }
            }];
        #endif
    } else if (@available(macOS 10.8, *)) {
        captureResult(channelID, true, NULL);
    } else {
        NSString *errorMsg = @"Notification API not available on this system";
        captureResult(channelID, false, [errorMsg UTF8String]);
    }
}

void checkNotificationAuthorization(int channelID) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            if (!ensureDelegateInitialized()) {
                NSString *errorMsg = @"Notification delegate has been lost. Reinitialize the notification service.";
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            [center getNotificationSettingsWithCompletionHandler:^(UNNotificationSettings *settings) {
                BOOL isAuthorized = (settings.authorizationStatus == UNAuthorizationStatusAuthorized);
                captureResult(channelID, isAuthorized, NULL);
            }];
        #endif
    } else if (@available(macOS 10.8, *)) {
        captureResult(channelID, true, NULL);
    } else {
        NSString *errorMsg = @"Notification API not available on this system";
        captureResult(channelID, false, [errorMsg UTF8String]);
    }
}

// Helper function to create notification content
UNMutableNotificationContent* createNotificationContent(const char *title, const char *subtitle, 
                                                       const char *body, const char *data_json, NSError **contentError) {
    NSString *nsTitle = [NSString stringWithUTF8String:title];
    NSString *nsSubtitle = subtitle ? [NSString stringWithUTF8String:subtitle] : @"";
    NSString *nsBody = [NSString stringWithUTF8String:body];
    
    UNMutableNotificationContent *content = [[UNMutableNotificationContent alloc] init];
    content.title = nsTitle;
    if (![nsSubtitle isEqualToString:@""]) {
        content.subtitle = nsSubtitle;
    }
    content.body = nsBody;
    content.sound = [UNNotificationSound defaultSound];
    
    // Parse JSON data if provided
    if (data_json) {
        NSString *dataJsonStr = [NSString stringWithUTF8String:data_json];
        NSData *jsonData = [dataJsonStr dataUsingEncoding:NSUTF8StringEncoding];
        NSError *error = nil;
        NSDictionary *parsedData = [NSJSONSerialization JSONObjectWithData:jsonData options:0 error:&error];
        if (!error && parsedData) {
            content.userInfo = parsedData;
        } else if (error) {
            *contentError = error;
        }
    }
    
    return content;
}

// Helper function to create legacy NSUserNotification content
NSUserNotification* createLegacyNotificationContent(const char *identifier, const char *title, 
                                                   const char *subtitle, const char *body,
                                                   const char *data_json, NSError **contentError) {
    NSUserNotification *notification = [[NSUserNotification alloc] init];
    notification.identifier = [NSString stringWithUTF8String:identifier];
    notification.title = [NSString stringWithUTF8String:title];
    
    if (subtitle && strlen(subtitle) > 0) {
        notification.subtitle = [NSString stringWithUTF8String:subtitle];
    }
    
    if (body && strlen(body) > 0) {
        notification.informativeText = [NSString stringWithUTF8String:body];
    }
    
    notification.soundName = NSUserNotificationDefaultSoundName;
    
    // Parse JSON data if provided
    if (data_json && strlen(data_json) > 0) {
        NSString *dataJsonStr = [NSString stringWithUTF8String:data_json];
        NSData *jsonData = [dataJsonStr dataUsingEncoding:NSUTF8StringEncoding];
        NSError *error = nil;
        NSDictionary *parsedData = [NSJSONSerialization JSONObjectWithData:jsonData options:0 error:&error];
        if (!error && parsedData) {
            notification.userInfo = parsedData;
        } else if (error) {
            *contentError = error;
        }
    }
    
    return notification;
}

void sendNotification(int channelID, const char *identifier, const char *title, const char *subtitle, const char *body, const char *data_json) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            if (!ensureDelegateInitialized()) {
                NSString *errorMsg = @"Notification delegate has been lost. Reinitialize the notification service.";
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];

            NSString *nsIdentifier = [NSString stringWithUTF8String:identifier];

            NSError *contentError = nil;
            UNMutableNotificationContent *content = createNotificationContent(title, subtitle, body, data_json, &contentError);
            
            if (contentError) {
                NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [contentError localizedDescription]];
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }

            UNTimeIntervalNotificationTrigger *trigger = nil;
            
            UNNotificationRequest *request = [UNNotificationRequest requestWithIdentifier:nsIdentifier content:content trigger:trigger];
            
            [center addNotificationRequest:request withCompletionHandler:^(NSError * _Nullable error) {
                if (error) {
                    NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [error localizedDescription]];
                    captureResult(channelID, false, [errorMsg UTF8String]);
                } else {
                    captureResult(channelID, true, NULL);
                }
            }];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            if (!ensureLegacyDelegateInitialized()) {
                NSString *errorMsg = @"Notification delegate has been lost. Reinitialize the notification service.";
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }

            NSError *contentError = nil;
            NSUserNotification *notification = createLegacyNotificationContent(identifier, title, subtitle, body, data_json, &contentError);
            
            [[NSUserNotificationCenter defaultUserNotificationCenter] deliverNotification:notification];
            captureResult(channelID, true, NULL);
        #endif
    } else {
        NSString *errorMsg = @"Notification API not available on this system";
        captureResult(channelID, false, [errorMsg UTF8String]);
    }
}

void sendNotificationWithActions(int channelID, const char *identifier, const char *title, const char *subtitle, 
                             const char *body, const char *categoryId, const char *data_json) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            if (!ensureDelegateInitialized()) {
                NSString *errorMsg = @"Notification delegate has been lost. Reinitialize the notification service.";
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];

            NSString *nsIdentifier = [NSString stringWithUTF8String:identifier];
            NSString *nsCategoryId = [NSString stringWithUTF8String:categoryId];
            
            NSError *contentError = nil;
            UNMutableNotificationContent *content = createNotificationContent(title, subtitle, body, data_json, &contentError);
            
            if (contentError) {
                NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [contentError localizedDescription]];
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }

            content.categoryIdentifier = nsCategoryId;
            
            UNTimeIntervalNotificationTrigger *trigger = nil;
            
            UNNotificationRequest *request = [UNNotificationRequest requestWithIdentifier:nsIdentifier content:content trigger:trigger];
            
            [center addNotificationRequest:request withCompletionHandler:^(NSError * _Nullable error) {
                if (error) {
                    NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [error localizedDescription]];
                    captureResult(channelID, false, [errorMsg UTF8String]);
                } else {
                    captureResult(channelID, true, NULL);
                }
            }];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            if (!ensureLegacyDelegateInitialized()) {
                NSString *errorMsg = @"Notification delegate has been lost. Reinitialize the notification service.";
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            NSError *contentError = nil;
            NSUserNotification *notification = createLegacyNotificationContent(identifier, title, subtitle, body, data_json, &contentError);
            
            if (contentError) {
                NSString *errorMsg = [NSString stringWithFormat:@"Error parsing data JSON: %@", [contentError localizedDescription]];
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            // Try to get the stored category
            if (categoryId && strlen(categoryId) > 0) {
                NSString *nsCategoryId = [NSString stringWithUTF8String:categoryId];
                NSDictionary *categoryData = getNotificationCategory(nsCategoryId);
                
                if (categoryData) {
                    NSArray *actions = categoryData[@"actions"];
                    
                    // NSUserNotification only supports one action button
                    if (actions.count > 0) {
                        NSDictionary *firstAction = actions.firstObject;
                        notification.hasActionButton = YES;
                        notification.actionButtonTitle = firstAction[@"title"];
                    }
                    
                    // If category has reply field, enable reply
                    if ([categoryData[@"hasReplyField"] boolValue]) {
                        notification.hasReplyButton = YES;
                    }
                    
                    // Store the category ID in the userInfo so we can reference it when handling responses
                    NSMutableDictionary *userInfo = notification.userInfo ? [notification.userInfo mutableCopy] : [NSMutableDictionary dictionary];
                    userInfo[@"categoryId"] = nsCategoryId;
                    notification.userInfo = userInfo;
                }
            }
            
            [[NSUserNotificationCenter defaultUserNotificationCenter] deliverNotification:notification];
            captureResult(channelID, true, NULL);
        #endif
    } else {
        NSString *errorMsg = @"Notification API not available on this system";
        captureResult(channelID, false, [errorMsg UTF8String]);
    }
}

void registerNotificationCategory(int channelID, const char *categoryId, const char *actions_json, bool hasReplyField, 
                                const char *replyPlaceholder, const char *replyButtonTitle) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            if (!ensureDelegateInitialized()) {
                NSString *errorMsg = @"Notification delegate has been lost. Reinitialize the notification service.";
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            NSString *nsCategoryId = [NSString stringWithUTF8String:categoryId];
            NSString *actionsJsonStr = actions_json ? [NSString stringWithUTF8String:actions_json] : @"[]";
            
            NSData *jsonData = [actionsJsonStr dataUsingEncoding:NSUTF8StringEncoding];
            NSError *error = nil;
            NSArray *actionsArray = [NSJSONSerialization JSONObjectWithData:jsonData options:0 error:&error];
            
            if (error) {
                NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [error localizedDescription]];
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            NSMutableArray *actions = [NSMutableArray array];
            
            for (NSDictionary *actionDict in actionsArray) {
                NSString *actionId = actionDict[@"id"];
                NSString *actionTitle = actionDict[@"title"];
                BOOL destructive = [actionDict[@"destructive"] boolValue];
                
                if (actionId && actionTitle) {
                    UNNotificationActionOptions options = UNNotificationActionOptionNone;
                    if (destructive) options |= UNNotificationActionOptionDestructive;
                    
                    UNNotificationAction *action = [UNNotificationAction 
                                                actionWithIdentifier:actionId
                                                title:actionTitle
                                                options:options];
                    [actions addObject:action];
                }
            }
            
            if (hasReplyField && replyPlaceholder && replyButtonTitle) {
                NSString *placeholder = [NSString stringWithUTF8String:replyPlaceholder];
                NSString *buttonTitle = [NSString stringWithUTF8String:replyButtonTitle];
                
                UNTextInputNotificationAction *textAction = 
                    [UNTextInputNotificationAction actionWithIdentifier:@"TEXT_REPLY"
                                                                title:buttonTitle
                                                            options:UNNotificationActionOptionNone
                                                textInputButtonTitle:buttonTitle
                                                textInputPlaceholder:placeholder];
                [actions addObject:textAction];
            }
            
            UNNotificationCategory *newCategory = [UNNotificationCategory 
                                            categoryWithIdentifier:nsCategoryId
                                            actions:actions
                                            intentIdentifiers:@[]
                                            options:UNNotificationCategoryOptionNone];
            
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            [center getNotificationCategoriesWithCompletionHandler:^(NSSet<UNNotificationCategory *> *categories) {
                NSMutableSet *updatedCategories = [NSMutableSet setWithSet:categories];
                
                // Remove existing category with same ID if it exists
                UNNotificationCategory *existingCategory = nil;
                for (UNNotificationCategory *category in updatedCategories) {
                    if ([category.identifier isEqualToString:nsCategoryId]) {
                        existingCategory = category;
                        break;
                    }
                }
                if (existingCategory) {
                    [updatedCategories removeObject:existingCategory];
                }
                
                [updatedCategories addObject:newCategory];
                [center setNotificationCategories:updatedCategories];

                captureResult(channelID, true, NULL);
            }];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            // Parse the actions to verify they're valid
            NSString *actionsJsonStr = actions_json ? [NSString stringWithUTF8String:actions_json] : @"[]";
            NSData *jsonData = [actionsJsonStr dataUsingEncoding:NSUTF8StringEncoding];
            NSError *error = nil;
            NSArray *actionsArray = [NSJSONSerialization JSONObjectWithData:jsonData options:0 error:&error];
            
            if (error) {
                NSString *errorMsg = [NSString stringWithFormat:@"Error: %@", [error localizedDescription]];
                captureResult(channelID, false, [errorMsg UTF8String]);
                return;
            }
            
            // Create a dictionary to store category information
            NSMutableDictionary *categoryData = [NSMutableDictionary dictionary];
            [categoryData setObject:actionsArray forKey:@"actions"];
            [categoryData setObject:@(hasReplyField) forKey:@"hasReplyField"];
            
            if (hasReplyField && replyPlaceholder) {
                [categoryData setObject:[NSString stringWithUTF8String:replyPlaceholder] forKey:@"replyPlaceholder"];
            }
            
            if (hasReplyField && replyButtonTitle) {
                [categoryData setObject:[NSString stringWithUTF8String:replyButtonTitle] forKey:@"replyButtonTitle"];
            }
            
            // Store the category
            NSString *nsCategoryId = [NSString stringWithUTF8String:categoryId];
            saveNotificationCategory(nsCategoryId, categoryData);
            
            captureResult(channelID, true, NULL);
        #endif
    } else {
        NSString *errorMsg = @"Notification API not available on this system";
        captureResult(channelID, false, [errorMsg UTF8String]);
    }
}

void removeNotificationCategory(int channelID, const char *categoryId) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            NSString *nsCategoryId = [NSString stringWithUTF8String:categoryId];
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            
            [center getNotificationCategoriesWithCompletionHandler:^(NSSet<UNNotificationCategory *> *categories) {
                NSMutableSet *updatedCategories = [NSMutableSet setWithSet:categories];
                
                // Find and remove the category with matching identifier
                UNNotificationCategory *categoryToRemove = nil;
                for (UNNotificationCategory *category in updatedCategories) {
                    if ([category.identifier isEqualToString:nsCategoryId]) {
                        categoryToRemove = category;
                        break;
                    }
                }
                
                if (categoryToRemove) {
                    [updatedCategories removeObject:categoryToRemove];
                    [center setNotificationCategories:updatedCategories];
                    captureResult(channelID, true, NULL);
                } else {
                    NSString *errorMsg = [NSString stringWithFormat:@"Category '%@' not found", nsCategoryId];
                    captureResult(channelID, false, [errorMsg UTF8String]);
                }
            }];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            // For legacy NSUserNotification, remove from our stored categories
            NSString *nsCategoryId = [NSString stringWithUTF8String:categoryId];
            NSDictionary *category = getNotificationCategory(nsCategoryId);
            
            if (category) {
                // Category exists, so remove it
                removeStoredNotificationCategory(nsCategoryId);
                captureResult(channelID, true, NULL);
            } else {
                // Category doesn't exist
                NSString *errorMsg = [NSString stringWithFormat:@"Category '%@' not found", nsCategoryId];
                captureResult(channelID, false, [errorMsg UTF8String]);
            }
        #endif
    } else {
        NSString *errorMsg = @"Notification API not available on this system";
        captureResult(channelID, false, [errorMsg UTF8String]);
    }
}

void removeAllPendingNotifications(void) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            [center removeAllPendingNotificationRequests];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            NSUserNotificationCenter *center = [NSUserNotificationCenter defaultUserNotificationCenter];
            [center removeAllDeliveredNotifications];
        #endif
    }
}

void removePendingNotification(const char *identifier) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            NSString *nsIdentifier = [NSString stringWithUTF8String:identifier];
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            [center removePendingNotificationRequestsWithIdentifiers:@[nsIdentifier]];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            NSString *nsIdentifier = [NSString stringWithUTF8String:identifier];
            NSUserNotificationCenter *center = [NSUserNotificationCenter defaultUserNotificationCenter];
            
            for (NSUserNotification *notification in center.deliveredNotifications) {
                if ([notification.identifier isEqualToString:nsIdentifier]) {
                    [center removeDeliveredNotification:notification];
                    break;
                }
            }
        #endif
    }
}

void removeAllDeliveredNotifications(void) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            [center removeAllDeliveredNotifications];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            NSUserNotificationCenter *center = [NSUserNotificationCenter defaultUserNotificationCenter];
            [center removeAllDeliveredNotifications];
        #endif
    }
}

void removeDeliveredNotification(const char *identifier) {
    if (@available(macOS 11.0, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 110000
            NSString *nsIdentifier = [NSString stringWithUTF8String:identifier];
            UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
            [center removeDeliveredNotificationsWithIdentifiers:@[nsIdentifier]];
        #endif
    } else if (@available(macOS 10.8, *)) {
        #if __MAC_OS_X_VERSION_MAX_ALLOWED >= 1080
            NSString *nsIdentifier = [NSString stringWithUTF8String:identifier];
            NSUserNotificationCenter *center = [NSUserNotificationCenter defaultUserNotificationCenter];
            
            for (NSUserNotification *notification in center.deliveredNotifications) {
                if ([notification.identifier isEqualToString:nsIdentifier]) {
                    [center removeDeliveredNotification:notification];
                    break;
                }
            }
        #endif
    }
}