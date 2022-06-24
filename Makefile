# generates the ent libraries
ent/generate:
	- go generate ./ent

# generates the required protobuf libraries
proto/generate:
	- rm -rfd gen
	- mkdir gen
	- buf generate buf.build/rltvty/ent-multiple-proto-bug-bar
	- buf generate buf.build/rltvty/ent-mutliple-proto-bug-foo
