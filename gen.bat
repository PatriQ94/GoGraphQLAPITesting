MOVE resolver.go resolver.old
go run github.com/99designs/gqlgen

COPY resolver.go resolver.old.new
git merge resolver.go resolver.go resolver.old