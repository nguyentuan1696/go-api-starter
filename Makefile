run:
	go run -race ./cmd/main.go serve

.PHONY: all deps deps-toolsaudit outdated vulncheck build debug watch-debug run watch-run lint lint-fix test watch-test weight coverage clean re