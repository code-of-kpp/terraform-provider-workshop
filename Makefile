DARK_THEME=catppuccin-mocha
LIGHT_THEME=catppuccin-latte

clean:
	./clean.sh

check-font-dark:
	presenterm check-font.md --theme=$(DARK_THEME)

check-font-light:
	presenterm check-font.md --theme=$(LIGHT_THEME)

talk-dark:
	presenterm README.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(DARK_THEME)

talk-light:
	presenterm README.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(LIGHT_THEME)

00-dark:
	presenterm 00-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(DARK_THEME)

00-light:
	presenterm 00-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(LIGHT_THEME)

01-dark:
	presenterm 01-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(DARK_THEME)

01-light:
	presenterm 01-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(LIGHT_THEME)

02-dark:
	presenterm 02-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(DARK_THEME)

02-light:
	presenterm 02-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(LIGHT_THEME)

03-dark:
	presenterm 03-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(DARK_THEME)

03-light:
	presenterm 03-*.md \
	  --enable-snippet-execution \
	  --enable-snippet-execution-replace \
	  --publish-speaker-notes \
	  --theme=$(LIGHT_THEME)
