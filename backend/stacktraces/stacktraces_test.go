package stacktraces

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openlyinc/pointy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzFormatStructureStackTrace(f *testing.F) {
	f.Fuzz(func(t *testing.T, stackTrace string) {
		output := FormatStructureStackTrace(context.TODO(), stackTrace)
		if stackTrace != "" && output == "" {
			t.Fatalf("expected to get an output")
		}
	})
}

func TestFormatStructureStackTrace(t *testing.T) {
	var inputs = []struct {
		language                      string
		stacktrace                    string
		expectedFrameError            string
		expectedFrameCount            int
		expectedInnerFrameCode        string
		expectedFramesWithFileNames   []bool
		expectedFramesWithLineNumbers []bool
	}{
		{language: "trpc", stacktrace: "Error: oh no!\n    at Procedure.resolve [as resolver] (webpack-internal:///(api)/./src/server/routers/name.ts:18:19)\n    at Array.<anonymous> (/Users/vkorolik/work/web-test/trpc-nextjs-demo-internal/node_modules/@trpc/server/dist/router-ee876044.cjs.dev.js:101:36)\n    at processTicksAndRejections (node:internal/process/task_queues:96:5)\n    at async callRecursive (/Users/vkorolik/work/web-test/trpc-nextjs-demo-internal/node_modules/@trpc/server/dist/router-ee876044.cjs.dev.js:119:24)\n    at async Procedure.call (/Users/vkorolik/work/web-test/trpc-nextjs-demo-internal/node_modules/@trpc/server/dist/router-ee876044.cjs.dev.js:144:20)\n    at async eval (webpack-internal:///(api)/./node_modules/@trpc/server/dist/resolveHTTPResponse-ab01e4b9.cjs.dev.js:205:24)\n    at async Promise.all (index 0)\n    at async Object.resolveHTTPResponse (webpack-internal:///(api)/./node_modules/@trpc/server/dist/resolveHTTPResponse-ab01e4b9.cjs.dev.js:201:24)\n    at async Object.nodeHTTPRequestHandler (webpack-internal:///(api)/./node_modules/@trpc/server/dist/nodeHTTPRequestHandler-9a93c255.cjs.dev.js:68:18)\n    at async eval (webpack-internal:///(api)/./node_modules/@trpc/server/adapters/next/dist/trpc-server-adapters-next.cjs.dev.js:48:5)", expectedFrameError: "Error: oh no!", expectedFrameCount: 10, expectedFramesWithFileNames: []bool{true, true, true, true, true, true, true, true, true, true}, expectedFramesWithLineNumbers: []bool{true, true, true, false, true, true, true, true, true, true}},
		{language: "python", stacktrace: "Traceback (most recent call last):\n  File \"/Users/vkorolik/Library/Caches/pypoetry/virtualenvs/highlight-io-T_znYNk9-py3.10/lib/python3.10/site-packages/flask/app.py\", line 2525, in wsgi_app\n    response = self.full_dispatch_request()\n  File \"/Users/vkorolik/Library/Caches/pypoetry/virtualenvs/highlight-io-T_znYNk9-py3.10/lib/python3.10/site-packages/flask/app.py\", line 1822, in full_dispatch_request\n    rv = self.handle_user_exception(e)\n  File \"/Users/vkorolik/Library/Caches/pypoetry/virtualenvs/highlight-io-T_znYNk9-py3.10/lib/python3.10/site-packages/flask/app.py\", line 1820, in full_dispatch_request\n    rv = self.dispatch_request()\n  File \"/Users/vkorolik/Library/Caches/pypoetry/virtualenvs/highlight-io-T_znYNk9-py3.10/lib/python3.10/site-packages/flask/app.py\", line 1796, in dispatch_request\n    return self.ensure_sync(self.view_functions[rule.endpoint])(**view_args)\n  File \"/Users/vkorolik/work/highlight/sdk/highlight-py/examples/app.py\", line 25, in hello\n    raise Exception(f\"random error! {idx}\")\nException: random error! 106\n", expectedFrameError: "Exception: random error! 106"},
		{language: "python-async", stacktrace: "Traceback (most recent call last):\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/middleware/base.py\", line 78, in call_next\n    message = await recv_stream.receive()\n              ^^^^^^^^^^^^^^^^^^^^^^^^^^^\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/anyio/streams/memory.py\", line 94, in receive\n    return self.receive_nowait()\n           ^^^^^^^^^^^^^^^^^^^^^\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/anyio/streams/memory.py\", line 87, in receive_nowait\n    raise EndOfStream\nanyio.EndOfStream\n\nDuring handling of the above exception, another exception occurred:\n\nTraceback (most recent call last):\n  File \"/home/vkorolik/work/highlight/sdk/highlight-py/highlight_io/sdk.py\", line 141, in trace\n    yield\n  File \"/home/vkorolik/work/highlight/sdk/highlight-py/highlight_io/integrations/fastapi.py\", line 22, in dispatch\n    return await call_next(request)\n           ^^^^^^^^^^^^^^^^^^^^^^^^\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/middleware/base.py\", line 84, in call_next\n    raise app_exc\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/middleware/base.py\", line 70, in coro\n    await self.app(scope, receive_or_disconnect, send_no_error)\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/middleware/exceptions.py\", line 79, in __call__\n    raise exc\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/middleware/exceptions.py\", line 68, in __call__\n    await self.app(scope, receive, sender)\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/fastapi/middleware/asyncexitstack.py\", line 21, in __call__\n    raise e\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/fastapi/middleware/asyncexitstack.py\", line 18, in __call__\n    await self.app(scope, receive, send)\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/routing.py\", line 706, in __call__\n    await route.handle(scope, receive, send)\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/routing.py\", line 276, in handle\n    await self.app(scope, receive, send)\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/starlette/routing.py\", line 66, in app\n    response = await func(request)\n               ^^^^^^^^^^^^^^^^^^^\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/fastapi/routing.py\", line 237, in app\n    raw_response = await run_endpoint_function(\n                   ^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n  File \"/home/vkorolik/.cache/pypoetry/virtualenvs/highlight-io-eYcfUyf6-py3.11/lib/python3.11/site-packages/fastapi/routing.py\", line 163, in run_endpoint_function\n    return await dependant.call(**values)\n           ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n  File \"/home/vkorolik/work/highlight/sdk/highlight-py/e2e/highlight_fastapi/main.py\", line 30, in root\n    logging.info(f'oh no {5 / 0}')\n                          ~~^~~\nZeroDivisionError: division by zero\n", expectedFrameError: "ZeroDivisionError: division by zero", expectedFrameCount: 17, expectedInnerFrameCode: "    logging.info(f'oh no {5 / 0}')"},
		{language: "golang", stacktrace: "\ngithub.com/highlight-run/highlight/backend/private-graph/graph.(*queryResolver).Admin\n\t/build/backend/private-graph/graph/schema.resolvers.go:6081\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query_admin.func2\n\t/build/backend/private-graph/graph/generated/generated.go:39227\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func4\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:72\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:110\ngithub.com/highlight/highlight/sdk/highlight-go.Tracer.InterceptField\n\t/build/sdk/highlight-go/tracer.go:47\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:109\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:110\ngithub.com/highlight-run/highlight/backend/util.Tracer.InterceptField\n\t/build/backend/util/tracer.go:45\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:109\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query_admin\n\t/build/backend/private-graph/graph/generated/generated.go:39225\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func280\n\t/build/backend/private-graph/graph/generated/generated.go:59279\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func3\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:69\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func281\n\t/build/backend/private-graph/graph/generated/generated.go:59284\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func282\n\t/build/backend/private-graph/graph/generated/generated.go:59288\ngithub.com/99designs/gqlgen/graphql.(*FieldSet).Dispatch.func1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/fieldset.go:42\nruntime.goexit\n\t/usr/local/go/src/runtime/asm_arm64.s:1172", expectedFrameError: "github.com/highlight-run/highlight/backend/private-graph/graph.(*queryResolver).Admin"},
		{language: "golang-panic", stacktrace: "\npanic: yo [recovered]\n\tpanic: yo\n\ngoroutine 22 [running]:\ntesting.tRunner.func1.2({0x103bc2120, 0x103de83a0})\n\t/usr/local/go/src/testing/testing.go:1396 +0x1c8\ntesting.tRunner.func1()\n\t/usr/local/go/src/testing/testing.go:1399 +0x378\npanic({0x103bc2120, 0x103de83a0})\n\t/usr/local/go/src/runtime/panic.go:884 +0x204\ngithub.com/highlight-run/highlight/backend/otel.structureStackTrace({0x10386edd2, 0x4d3})\n\t/Users/vkorolik/work/highlight/backend/otel/parse.go:48 +0x270\ngithub.com/highlight-run/highlight/backend/otel.TestFormatStructureStackTrace.func1(0x0?)\n\t/Users/vkorolik/work/highlight/backend/otel/parse_test.go:27 +0x38\ntesting.tRunner(0x14000a4e000, 0x14000a205b0)\n\t/usr/local/go/src/testing/testing.go:1446 +0x10c\ncreated by testing.(*T).Run\n\t/usr/local/go/src/testing/testing.go:1493 +0x300\n", expectedFrameError: "panic: yo [recovered]"},
		{language: "golang-panic-recover", stacktrace: "\ngithub.com/highlight/highlight/sdk/highlight-go.GraphQLRecoverFunc.func1\n\t/home/vkorolik/work/highlight/sdk/highlight-go/tracer.go:110\ngithub.com/99designs/gqlgen/graphql.(*OperationContext).Recover\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/context_operation.go:124\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query_error_instance.func1\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:40988\nruntime.gopanic\n\t/usr/local/go/src/runtime/panic.go:890\ngithub.com/highlight-run/highlight/backend/private-graph/graph.(*Resolver).isAdminInProject\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/resolver.go:481\ngithub.com/highlight-run/highlight/backend/private-graph/graph.(*Resolver).isAdminInProjectOrDemoProject\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/resolver.go:317\ngithub.com/highlight-run/highlight/backend/private-graph/graph.(*Resolver).doesAdminOwnErrorGroup\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/resolver.go:875\ngithub.com/highlight-run/highlight/backend/private-graph/graph.(*Resolver).canAdminViewErrorGroup\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/resolver.go:911\ngithub.com/highlight-run/highlight/backend/private-graph/graph.(*queryResolver).ErrorInstance\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/schema.resolvers.go:4097\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query_error_instance.func2\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:40994\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func4\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:72\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8.1\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:110\ngithub.com/highlight/highlight/sdk/highlight-go.Tracer.InterceptField\n\t/home/vkorolik/work/highlight/sdk/highlight-go/tracer.go:59\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:109\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8.1\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:110\ngithub.com/highlight-run/highlight/backend/util.Tracer.InterceptField\n\t/home/vkorolik/work/highlight/backend/util/tracer.go:45\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:109\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query_error_instance\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:40992\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func43\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:67991\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func3\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:69\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func44\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:67996\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func45\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:68000\ngithub.com/99designs/gqlgen/graphql.(*FieldSet).Dispatch\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/fieldset.go:34\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:70398\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executableSchema).Exec.func1\n\t/home/vkorolik/work/highlight/backend/private-graph/graph/generated/generated.go:8576\ngithub.com/99designs/gqlgen/graphql/executor.(*Executor).DispatchOperation.func1.1.1\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/executor.go:119\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func2\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:66\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func6.1\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:92\ngithub.com/highlight/highlight/sdk/highlight-go.Tracer.InterceptResponse\n\t/home/vkorolik/work/highlight/sdk/highlight-go/tracer.go:93\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func6\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:91\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func6.1\n\t/home/vkorolik/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:92\ngithub.com/highlight-run/highlight/backend/util.Tracer.InterceptResponse\n\t/home/vkorolik/work/highlight/backend/util/tracer.go:67", expectedFrameError: "github.com/highlight-run/highlight/backend/private-graph/graph.(*Resolver).isAdminInProject"},
		{language: "golang-extended", stacktrace: "\ngithub.com/highlight-run/highlight/backend/private-graph/graph.(*queryResolver).SubscriptionDetails\n\t/build/backend/private-graph/graph/schema.resolvers.go:6520\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query_subscription_details.func2\n\t/build/backend/private-graph/graph/generated/generated.go:41892\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func4\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:72\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:110\ngithub.com/highlight/highlight/sdk/highlight-go.Tracer.InterceptField\n\t/build/sdk/highlight-go/tracer.go:47\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:109\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:110\ngithub.com/highlight-run/highlight/backend/util.Tracer.InterceptField\n\t/build/backend/util/tracer.go:45\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func8\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:109\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query_subscription_details\n\t/build/backend/private-graph/graph/generated/generated.go:41890\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func310\n\t/build/backend/private-graph/graph/generated/generated.go:62598\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func3\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:69\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func311\n\t/build/backend/private-graph/graph/generated/generated.go:62603\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query.func312\n\t/build/backend/private-graph/graph/generated/generated.go:62607\ngithub.com/99designs/gqlgen/graphql.(*FieldSet).Dispatch\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/fieldset.go:34\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executionContext)._Query\n\t/build/backend/private-graph/graph/generated/generated.go:63005\ngithub.com/highlight-run/highlight/backend/private-graph/graph/generated.(*executableSchema).Exec.func1\n\t/build/backend/private-graph/graph/generated/generated.go:7703\ngithub.com/99designs/gqlgen/graphql/executor.(*Executor).DispatchOperation.func1.1.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/executor.go:119\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func2\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:66\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func6.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:92\ngithub.com/highlight/highlight/sdk/highlight-go.Tracer.InterceptResponse\n\t/build/sdk/highlight-go/tracer.go:75\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func6\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:91\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func6.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:92\ngithub.com/highlight-run/highlight/backend/util.Tracer.InterceptResponse\n\t/build/backend/util/tracer.go:65\ngithub.com/99designs/gqlgen/graphql/executor.processExtensions.func6\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/extensions.go:91\ngithub.com/99designs/gqlgen/graphql/executor.(*Executor).DispatchOperation.func1.1\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/executor/executor.go:118\ngithub.com/99designs/gqlgen/graphql/handler/transport.POST.Do\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/handler/transport/http_post.go:89\ngithub.com/99designs/gqlgen/graphql/handler.(*Server).ServeHTTP\n\t/go/pkg/mod/github.com/99designs/gqlgen@v0.17.24/graphql/handler/server.go:121\ngithub.com/go-chi/chi.(*Mux).routeHTTP\n\t/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go:431\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2122\ngithub.com/highlight/highlight/sdk/highlight-go/middleware/chi.Middleware.func1\n\t/build/sdk/highlight-go/middleware/chi/middleware.go:20\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2122", expectedFrameError: "github.com/highlight-run/highlight/backend/private-graph/graph.(*queryResolver).SubscriptionDetails"},
		{language: "next.js-backend", stacktrace: "Error: GraphQL Error (Code: 401): {\"response\":{\"error\":\"{\\\"errors\\\":[{\\\"message\\\":\\\"token verification failed: token contains an invalid number of segments\\\"}],\\\"data\\\":null}\",\"status\":401,\"headers\":{}},\"request\":{\"query\":\"\\n      query GetPosts() {\\n        posts(orderBy: publishedAt_DESC) {\\n          slug\\n        }\\n      }\\n    \"}}\n    at /Users/jaykhatri/projects/highlight.io/node_modules/graphql-request/dist/index.js:416:31\n    at step (/Users/jaykhatri/projects/highlight.io/node_modules/graphql-request/dist/index.js:67:23)\n    at Object.next (/Users/jaykhatri/projects/highlight.io/node_modules/graphql-request/dist/index.js:48:53)\n    at fulfilled (/Users/jaykhatri/projects/highlight.io/node_modules/graphql-request/dist/index.js:39:58)\n    at process.processTicksAndRejections (node:internal/process/task_queues:95:5)", expectedFrameError: "Error: GraphQL Error (Code: 401): {\"response\":{\"error\":\"{\\\"errors\\\":[{\\\"message\\\":\\\"token verification failed: token contains an invalid number of segments\\\"}],\\\"data\\\":null}\",\"status\":401,\"headers\":{}},\"request\":{\"query\":\"\\n      query GetPosts() {\\n        posts(orderBy: publishedAt_DESC) {\\n          slug\\n        }\\n      }\\n    \"}}"},
		{language: "node.js-console", stacktrace: "\"Error\\n    at console.<computed> [as error] (webpack-internal:///(api)/../../sdk/highlight-node/dist/index.mjs:194:15)\\n    at DevServer.logErrorWithOriginalStack (/Users/vkorolik/work/highlight/e2e/nextjs/node_modules/next/dist/server/dev/next-dev-server.js:803:71)\\n    at processTicksAndRejections (node:internal/process/task_queues:96:5)\"", expectedFrameError: "Error"},
		{language: ".NET", stacktrace: "System.Exception: oh no, a random error occurred 1a77a6d6-4803-4de8-822b-13a62397b9d3\n   at Program.<>c__DisplayClass0_0.<<Main>$>b__2() in /home/vkorolik/work/highlight/e2e/dotnet/Program.cs:line 89\n   at lambda_method3(Closure, Object, HttpContext)\n   at Microsoft.AspNetCore.HttpsPolicy.HttpsRedirectionMiddleware.Invoke(HttpContext context) in /_/src/aspnetcore/artifacts/source-build/self/src/src/Middleware/HttpsPolicy/src/HttpsRedirectionMiddleware.cs:line 88\n   at Microsoft.AspNetCore.StaticFiles.StaticFileMiddleware.Invoke(HttpContext context) in /_/src/aspnetcore/artifacts/source-build/self/src/src/Middleware/StaticFiles/src/StaticFileMiddleware.cs:line 82\n   at Swashbuckle.AspNetCore.SwaggerUI.SwaggerUIMiddleware.Invoke(HttpContext httpContext)\n   at Swashbuckle.AspNetCore.Swagger.SwaggerMiddleware.Invoke(HttpContext httpContext, ISwaggerProvider swaggerProvider)\n   at Microsoft.AspNetCore.Diagnostics.DeveloperExceptionPageMiddlewareImpl.Invoke(HttpContext context) in /_/src/aspnetcore/artifacts/source", expectedFrameError: "oh no, a random error occurred 1a77a6d6-4803-4de8-822b-13a62397b9d3", expectedFrameCount: 7, expectedFramesWithFileNames: []bool{true, false, true, true, false, false, true}, expectedFramesWithLineNumbers: []bool{true, true, true, true, true, true, false}},
	}
	for _, input := range inputs {
		t.Run(input.language, func(t *testing.T) {
			frames, err := StructureOTELStackTrace(input.stacktrace)
			str, _ := json.MarshalIndent(frames, "", "\t")
			fmt.Printf("%s\n", str)
			assert.NoErrorf(t, err, "unexpected error")
			if input.expectedFrameCount != 0 {
				assert.Equal(t, input.expectedFrameCount, len(frames))
			}
			if input.expectedInnerFrameCode != "" {
				assert.Equal(t, input.expectedInnerFrameCode, *frames[0].LineContent)
			}
			for idx, frame := range frames {
				assert.NotNil(t, frame)
				assert.NotNil(t, frame.FileName)
				if len(input.expectedFramesWithFileNames) > 0 {
					if input.expectedFramesWithFileNames[idx] {
						assert.GreaterOrEqual(t, len(pointy.StringValue(frame.FileName, "")), 1, idx)
						if input.expectedFramesWithLineNumbers[idx] {
							assert.GreaterOrEqual(t, pointy.IntValue(frame.LineNumber, 0), 1, idx)
						} else {
							assert.LessOrEqual(t, pointy.IntValue(frame.LineNumber, 0), 0, idx)
						}
					} else {
						assert.LessOrEqual(t, len(pointy.StringValue(frame.FileName, "")), 0, idx)
					}
				} else {
					assert.GreaterOrEqual(t, len(pointy.StringValue(frame.FileName, "")), 1, idx)
					assert.GreaterOrEqual(t, pointy.IntValue(frame.LineNumber, 0), 1, idx)
				}
				assert.Equal(t, input.expectedFrameError, *frame.Error)
			}
		})
	}
}
