var gulp = require("gulp");
var shell = require("gulp-shell");

// This compiles new binaries with source changes
gulp.task("install-binary", shell.task([
    'go install'
]));

// Second argument tells install-binary is a dependency for restart-supervisor
gulp.task("restart-supervisor", shell.task([
    'supervisorctl restart myserver'
]))

// Watch the source code for all changes
gulp.task('default', function() {
    gulp.watch("*", gulp.series("install-binary", "restart-supervisor"));
});