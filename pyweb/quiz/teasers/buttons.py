# Push My Buttons

# Create UI
display = []
buttons = []
for n in range(10):
    # A button is a function called when user clicks on it
    buttons.append(lambda: display.append(n))

# Click on button 3
btn = buttons[3]
btn()
print(display)
