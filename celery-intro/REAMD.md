* The propoganda from http://celeryproject.org/
* "Celery is ..." and "Features" from 
  http://docs.celeryproject.org/en/latest/getting-started/introduction.html
* http://docs.celeryproject.org/en/latest/getting-started/first-steps-with-celery.html
* Open `tasks.py`
* Go over `celery = Celery(...` configuration
    * Show http://docs.celeryproject.org/en/latest/configuration.html#configuration
    * Error email
* celery -A tasks worker --loglevel=info
    * Show outpot
* Go over `add` task
    * task is a function
    * some limit on type of variables (depends on serializer)
    * decorator
    * Show you can use OO (http://docs.celeryproject.org/en/latest/reference/celery.app.task.html)
    * task options
    * http://docs.celeryproject.org/en/latest/userguide/tasks.html
* Open ipy/notebook?
* Calling
    * import tasks
    * res = tasks.add(4, 4)  # Show it's not running
    * res = tasks.add.delay(4, 4)  # AsyncResult
        * res.id, res.ready()
    * Show log from master
    * res.get()
* Kill server, run task, restart server
    * Show it picks up the task
    * Reloading with `kill -HUP` (does not work on OSX) 
        - http://docs.celeryproject.org/en/latest/userguide/workers.html#pool-restart-command
    * Also celery.control module (http://docs.celeryproject.org/en/latest/userguide/workers.html)
* Task properties
    * http://docs.celeryproject.org/en/latest/userguide/tasks.html#general
* Canvas
    * http://docs.celeryproject.org/en/latest/userguide/canvas.html
* Periodic tasks
    * http://docs.celeryproject.org/en/latest/userguide/periodic-tasks.html
    * http://docs.celeryproject.org/en/latest/faq.html#module-celery.task.base
* Go over user guide
    * http://docs.celeryproject.org/en/latest/userguide/index.html
    * Monitoring
    * Debugging - http://docs.celeryproject.org/en/latest/tutorials/debugging.html
        * Show with mul (print args + show mul.request)
        * Security
* Extras
    * Flower
        * http://localhost:5555
        * Has ReST API
    * celery events
    * celerymon
    * django-celery
