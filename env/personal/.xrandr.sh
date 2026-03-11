#!/bin/sh

if xrandr | grep -q "^DP-0 connected" && xrandr | grep -q "^HDMI-0 connected"; then
    xrandr --output HDMI-0 --mode 1920x1080 --pos 0x360 --rotate normal \
           --output DP-0 --primary --mode 2560x1440 --pos 1920x0 --rotate normal \
           --output DP-1 --off --output DP-2 --off --output DP-3 --off --output DP-4 --off --output DP-5 --off
fi
