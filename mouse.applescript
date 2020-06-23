try
	tell application "System Preferences"
		(run)
		set current pane to pane "com.apple.preference.trackpad"
	end tell
	tell application "System Events"
		tell process "System Preferences"
			delay 0.6
			click radio button "Scroll & Zoom" of tab group 1 of window "Trackpad"
			click checkbox 1 of tab group 1 of window "Trackpad"
		end tell
		
		tell application "System Preferences" to quit
	end tell
end try