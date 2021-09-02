# Status.io subscriber upload

A simple cli tool to bulk upload a list of email subscribers from a file.

The file format is 1 email address per line.  eg:

```
foo@example.com
bar@example.org
baz@example.net
```

It does _not_ handle the common approach of adding extra info on the line. eg:

```
"Foo Nurk" <foo@example.com>
"Bar Nurk" <bar@example.org>
"Baz Nurk" <baz@example.net>
```

So, keep it simple.

## Usage

To use this utility, set environment variables for the main status page values:

* API_ID
* API_KEY
* STATUS_PAGE_ID

Then call this utility with a filename as the first command line argument. eg:

```
$ API_ID=abc API_KEY=def STATUS_PAGE_ID=ghi ./statusio_metrics_updater example_subscriber_list.txt
```
