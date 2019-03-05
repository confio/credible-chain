#!/bin/bash

HOST=${HOST:-localhost}

curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":2,"rep_vote":"AL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"axx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'
curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":3,"rep_vote":"BL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"bxx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'
curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":2,"rep_vote":"CL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"cxx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'
curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":3,"rep_vote":"DL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"dxx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'
curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":2,"rep_vote":"EL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"exx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'
curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":3,"rep_vote":"FL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"fxx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'
curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":2,"rep_vote":"GL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"gxx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'
curl -X POST http://$HOST:5005/vote --data '{"vote":{"main_vote":3,"rep_vote":"HL1","charity":"MD","postCode":"SW16","birth_year":1980,"donation":100},"identifier":"hxx123456xxz","sms_code":"BA383SKD 10","voted_at":"2019-02-21T13:27:20.070566467Z"}'

