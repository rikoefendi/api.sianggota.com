dev:
	#Watching File
	reflex -r '\.go' -s -- sh -c "ENV=dev.env go run app/main.go"
commit:
	git add .
	git commit -m "$(message)"
	git push origin $(origin)
push:
	git push origin $(origin)
pull:
	git pull origin $(origin)