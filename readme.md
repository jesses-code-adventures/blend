![Tests](https://github.com/jesses-code-adventures/blend/actions/workflows/test.yml/badge.svg)

# blend

llm tool that ingests codesbases, applies diffs, runs the code.

## design

![Design](docs/design.png)

## dev

### setup

clone the repo

```bash
git clone https://github.com/jesses-code-adventures/blend
```

create a `.env.mine` file and add your `OPENAI_API_KEY` and `TEST_PRECEDENCE_VAR=.env.mine` (this will be removed at some point).

```.env.mine
OPENAI_API_KEY=sk_thisisanapikey_4000
TEST_PRECEDENCE_VAR=.env.mine
```

run `make stream-test` to check the openai integration is working.

### environment vars

inside the go environment, env var files are loaded in the order of .env.public, then .env.mine with override behaviour, so you can clone the repo and use the .env.public variables for local dev and customize them in .env.mine

.env.test is also available and committed for env vars that should only be used in tests.

check all your env vars by running `make dump`.

## todo

- [x] openai integration
- [ ] claude integration
- [ ] llama integration
- [ ] generic implementation for locally hosted llms
- [x] llm responds with unstructured response
- [x] llm responds with code-only response
- [ ] parse llm response as git diffs comparing to original files
- [ ] allow user to accept full diff and output to different file
- [ ] allow user to accept full diff and output to original file (overwrite the input inplace)
- [ ] allow user to step through diffs and accept/retry
- [ ] allow llm to apply diffs directly to original file
- [ ] allow llm to run overwritten code and restart the loop if it fails to execute
- [ ] allow llm to run overwritten code and restart the loop if tests fail
