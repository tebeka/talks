# Effective Serialization

## Description

> Both your title and this description are made public and displayed in the
> conference program to help attendees decide whether they are interested in
> this presentation. Limit this description to a few concise paragraphs.

At the "edges" of your program (e.g. when interacting with other services) you will use serialization to send and receive data.
Working efficiently with serialization will make your interfaces more flexible, will save you money on CPU & bandwidth and free you to focus on writing business value code.

In this tutorial we'll cover best practices of serialization and work with two very popular serialization formats: JSON and protocol buffers.


## Audience

> 1–2 paragraphs that should answer three questions: (1) Whom is this tutorial
> for? (2) What background, knowledge, or experience do you expect attendees to
> have? (3) What do you expect attendees to learn, or to be able to do after
> attending your tutorial?

The audience are Python developers who like to become more proficient with serialization.
They audience know Python and be familiar with the command line.
At the end of the attendees will be able to work with JSON & protocol buffers, serialize custom types and avoid common pitfalls.

## Level

Beginner

## Format

> We’d like to learn how your tutorial will be structured. Aside from lecture,
> will attendees be engaging with the material in other ways? If there is any
> particular way you'd like to interact with remote attendees or guide them
> through the tutorial, you may describe it here as well. (e.g. chat, quizzes,
> polls, iPython notebooks, etc).

I'll start with a lecture (~20min) about serialization in general, discuss how to pick a serialization format and go over some of the pros and cons of various serialization format.

Then we'll start developing an application that stores and serves weather data, working first with JSON and then with protocol buffers. At every step I'll describe the current problem, we'll discuss the solution and then the attendees will have time to implement before I show my solution.

The attendees will need Python and the [protoc](https://developers.google.com/protocol-buffers/docs/downloads) compiler installed and have internet connection.

## Outline

> Make an outline that lists the topics and activities you will guide your
> students through during your 3 hour tutorial.
> 
> Please estimate what portion of your tutorial you’ll devote to each major
> topic (usually there are 2-5 of those). It’s fine if these timings change
> when you give the tutorial. We just want to see where the focus will be.

- Serialization overview (20min)
- Application overview & data types (10min)
- Basic application with JSON (40min)
- Custom types with JSON (45min)
- Protocol buffers (50min)
- Benchmarking CPU & memory (10min)

## Past Experience

> Please summarize your teaching or public speaking experience, as well as your
> experience with the subject of the tutorial. If you have offered this
> tutorial before, please provide links to the material and video, if possible.
> Otherwise, please provide links to one (or two!) previous presentations by
> each speaker. If you have any additional notes, they can be placed here as
> well.

I'm an experienced Python (23+ years) developer and have extensive speaking experience.
I'm using serialization in every project I write and have experience both with JSON, protocol buffers and many other serialization formats.

I gave a Go version of this tutorial, see links below.

### Previous Talks & Blogs
- [IPython: The Productivity Booster](https://www.youtube.com/watch?v=Zmo2ZN1SJ_Q) from EuroPython
- [ArdanLabs YouTube Channel](https://www.youtube.com/channel/UCCgGRKeRM1b0LTDqqb4NqjA/videos) English, mostly Go content
- [Descriptors - Supercharge Your Dot](https://www.youtube.com/watch?v=J2jI-mSyG74&t=4s) at PyCon Israel (Hebrew)
- [All package managers suck, conda suck less](https://www.youtube.com/watch?v=ywsf3mmv6do) at PyCon Israel (Hebrew)
- [Writing Command Line Friendly Applications](https://www.youtube.com/watch?v=bbaypngQ7mE&t=1s) at PyCon Israel (Hebrew)
- [Working with SQLite using Go & Python](https://www.ardanlabs.com/blog/2020/11/working-with-sqlite-using-go-python.html) blog post
