# Student Info

A simple example web app for querying student enrollment info from a CSV database that runs as a standalone binary with assets embeded, frontend code in HTMX and backend using Go templates.

To run it type:

```bash
go run .
```

and open [localhost:8080](http://localhost:8080) in a web browser.  If you would like to start it on a different port, set the `PORT` environment variable, for example:

```bash
PORT=3000 go run .
```

You can modify the student data by editing the data.csv file in a text editor or spreadsheet software and then turning the studentinfo server on and off again.  The order of the columns is:

1. ID
2. Name
3. Instrument
4. Teacher

If you compile the code, everything that's needed for the web server (except the data.csv file) will be embeded into the binary so that you only need to copy those 2 things, e.g. on Windows:

```bash
go build -O studentinfo.exe .
```

Now you can copy data.csv and studentinfo.exe to a server that can run Windows programs and start it without needing to install anything else there.  Just like locally, you can edit the CSV file and restart the server to make changes.

---

Written by Siôn le Roux, for licence details see the LICENSE file.
