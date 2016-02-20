var gulp = require('gulp');
var path = require('path');
var shell = require('gulp-shell');

var goPath = 'src/main/**/*.go';

var template = {
        templateData: {
            stripPath: function(filePath) {
                var subPath = filePath.substring(process.cwd().length + 5);
                var pkg = subPath.substring(0, subPath.lastIndexOf(path.sep));
                return pkg;
            }
        }
    };

gulp.task('compilepkg', function() {
    return gulp.src(goPath, {read: false})
        .pipe(shell(['go install <%= stripPath(file.path) %>'], template));
});

gulp.task('fmt', function() {
    return gulp.src(goPath, {read: false})
        .pipe(shell(['go fmt <%= stripPath(file.path) %>'], template));
});

gulp.task('watch', function() {
    gulp.watch(goPath, ['compilepkg', 'fmt']);
});
