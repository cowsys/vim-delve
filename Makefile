default: test

.PHONY: test
test:
	cd test; PROFILE_LOG=profile.txt themis --reporter dot *.vimspec

.PHONY: gen-coverage
gen-coverage:
	cd test; covimerage write_coverage profile.txt

cover: test gen-coverage
	cd test; coverage report
