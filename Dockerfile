FROM ubuntu:latest\nRUN apt-get update && apt-get install -y curl git\nWORKDIR /app\nCOPY . /app\nCMD ["./your-script"]
