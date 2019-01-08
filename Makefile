hashnode:
	go install github.com/hashnode/hashnode-cli
dry-release:
	goreleaser release --skip-publish --snapshot --rm-dist
