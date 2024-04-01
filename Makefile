default: setup_chapter

setup_chapter:
	@mkdir -p $(CHAPTER) && cd $(CHAPTER) && go mod init $(CHAPTER) && touch main.go

run:
	@go run $(CHAPTER)/main.go
