#import "badge_darwin.h"
#import <Cocoa/Cocoa.h>

void setBadgeLabel(const char *label) {
    NSString *nsLabel = nil;
	if (label != NULL) {
		nsLabel = [NSString stringWithUTF8String:label];
	}
	[[NSApp dockTile] setBadgeLabel:nsLabel];
	[[NSApp dockTile] display];
}