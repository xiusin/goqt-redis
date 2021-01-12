#!/usr/local/bin/v run
defer {
	rm("build")
}
println ('构建windows_64_static')
exec("qtdeploy -docker build windows_64_static") or {
	panic(err)
}

