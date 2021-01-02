# Embedding Python in Your Application

## Description

> Both your title and this description are made public and displayed in the
> conference program to help attendees decide whether they are interested in
> this presentation. Limit this description to a few concise paragraphs.

Most people use Python's C API to write extension modules. But the C API can also be used to embed Python.

Embedding Python in your application allows users to extend it without requiring your help.  In performance sensitive applications, you can call Python code from other languages without the overhead of serialization and network calls.

In this talk we'll discuss how to embed Python and will see an example of Go code calling Python/numpy code with 0 serialization and ε memory allocations.

## Who and Why (Audience)

> 1–2 paragraphs that should answer three questions: (1) Who is this talk for?
> (2) What background knowledge or experience do you expect the audience to
> have? (3) What do you expect the audience to learn or do after watching the
> talk?

The audience are Python (and non-Python) developers who'd like to extend an application with Python code. The audience should know Python & a little bit of C.
At the end of the talk, the audience should know how to embed Python in their application.

## Level

Intermediate


## Outline

> The “outline” is a skeleton of your talk that is as detailed as possible,
> including rough timings or estimates for different sections. If requesting a
> 45 minute slot, please describe what content would appear in the 45 minute
> version but not a 30 minute version, either within the outline or in a
> paragraph at the end.

- The Python C API (5min)
- Embedding use cases (5min)
- A simple embedding example (5min)
- Embedding Python in a Go application (15min)

## Additional Notes

> Anything else you would like to share with the committee: Speaker public
> speaking experience. Speaker subject matter experience. Have the speaker(s)
> given this presentation before elsewhere? Links to recordings, slides, blog
> posts, code, or other material. Specific needs or special requests —
> accessibility, audio (will you need to play pre-recorded sound?), or
> restrictions on when your talk can be scheduled.

I'm an experienced Python (23+ years) and Go (11+ years) developer and have extensive speaking experience.

I have embedded Python in a couple of applications previously: A hardware simulator and a serverless platform (https://nuclio.io/).

This talk will be based on [this blog post](https://www.ardanlabs.com/blog/2020/09/using-python-memory.html) but will be adapted to Python developers.

### Previous Talks & Blogs
- [IPython: The Productivity Booster](https://www.youtube.com/watch?v=Zmo2ZN1SJ_Q) from EuroPython
- [ArdanLabs YouTube Channel](https://www.youtube.com/channel/UCCgGRKeRM1b0LTDqqb4NqjA/videos) English, mostly Go content
- [Descriptors - Supercharge Your Dot](https://www.youtube.com/watch?v=J2jI-mSyG74&t=4s) at PyCon Israel (Hebrew)
- [All package managers suck, conda suck less](https://www.youtube.com/watch?v=ywsf3mmv6do) at PyCon Israel (Hebrew)
- [Writing Command Line Friendly Applications](https://www.youtube.com/watch?v=bbaypngQ7mE&t=1s) at PyCon Israel (Hebrew)
- [Working with SQLite using Go & Python](https://www.ardanlabs.com/blog/2020/11/working-with-sqlite-using-go-python.html) blog post
