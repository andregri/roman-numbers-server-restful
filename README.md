# For developers
You can install `supervisord` and `gulp` for a easier development.

The file `supervisord/goproject.conf` has to be copied to `/etc/supervisor/conf.d/goproject.conf`.

# Tests

```
curl -X GET "http://localhost:8000/roman_number/5"
curl -X GET "http://localhost:8000/roman_number/12"
curl -X GET "http://localhost:8000/random_resource/10"
```
