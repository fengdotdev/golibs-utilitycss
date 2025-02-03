VERSION = 0.0.8

updatev:
		git tag v${VERSION} && git push origin v${VERSION}
