.DEFAULT_GOAL := help
.PHONY: help start

## Start server
start:
	go run bot-currency.go -web

help:
	$(info Available targets)
	@awk '/^[a-zA-Z\-\_0-9]+:/ {                                   \
	  nb = sub( /^## /, "", helpMsg );                             \
	  if(nb == 0) {                                                \
	    helpMsg = $$0;                                             \
	    nb = sub( /^[^:]*:.* ## /, "", helpMsg );                  \
	  }                                                            \
	  if (nb)                                                      \
	    printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg; \
	}                                                              \
	{ helpMsg = $$0 }'                                             \
	width=$$(grep -o '^[a-zA-Z_0-9]\+:' $(MAKEFILE_LIST) | wc -L)  \
	$(MAKEFILE_LIST)
