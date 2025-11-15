format:
	goimports -w ./ies/*
	gofmt -w ./ies/*
	echo "✓ Formatting complete!"