//go:build darwin

#ifndef BADGE_DARWIN_H
#define BADGE_DARWIN_H

#import <Foundation/Foundation.h>

void setBadgeLabel(const char *label);

#endif /* BADGE_DARWIN_H */