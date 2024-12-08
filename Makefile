# Makefile

# Directories
ASSETS_DIR := ./assets
BUNDLED_DIR := ./bundled

# Files to generate
ANIMATED_BUNDLE := $(BUNDLED_DIR)/animated.go
FONT_BUNDLE := $(BUNDLED_DIR)/font.go

# Default target
bundle: $(ANIMATED_BUNDLE) $(FONT_BUNDLE) fix-packages

# Bundle animated assets
$(ANIMATED_BUNDLE): $(ASSETS_DIR)/animated
	mkdir -p $(BUNDLED_DIR)
	fyne bundle -o $@ $<

# Bundle font assets
$(FONT_BUNDLE): $(ASSETS_DIR)/font
	mkdir -p $(BUNDLED_DIR)
	fyne bundle -o $@ $<

# Fix package declaration and variable names
fix-packages: $(ANIMATED_BUNDLE) $(FONT_BUNDLE)
	@for file in $(BUNDLED_DIR)/*.go; do \
		sed -i 's/package main/package bundled/' $$file; \
		sed -i 's/var resource/var Resource/' $$file; \
	done

# Clean generated files
clean:
	rm -rf $(BUNDLED_DIR)

# Phony targets
.PHONY: bundle fix-packages clean
