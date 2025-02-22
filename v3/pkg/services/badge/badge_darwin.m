#import "badge_darwin.h"
#import <Cocoa/Cocoa.h>

void setBadgeLabel(const char *label) {
    @autoreleasepool {
        NSString *nsLabel = nil;
        if (label != NULL) {
            nsLabel = [NSString stringWithUTF8String:label];
        }
        dispatch_async(dispatch_get_main_queue(), ^{
            [[NSApp dockTile] setBadgeLabel:nsLabel];
            [[NSApp dockTile] display];
        });
    }
}