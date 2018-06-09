# Data Management - We can do better

Talk given at [Big Data Analytics meetup](https://www.meetup.com/Big-Data-Analytics-Israel/events/251392450/) June 18, 2018

## Ideas
* Regular schema not enough (inch vs feet ...)
    * Weather example where in c/10
* Column names in database
* Ask TT about data quality regulation
    - GDPR forces companies
    - error rate KPI (if above threshold, reject)
    - Header (# of records)
    - Compare to previous period (anomaly)
* Track source of data through ETL
* Bit rot
    * Example of RAM cosmic rays
    * SSD on shelf
* Monitoring & Alerting
* Dependencies between values (height & weight)
* GIGO
* Story about missing publisher ID (two weeks of data?)
* Story about reverse proxy: 2 month work for not understanding the ETL
* KPIs for data quality
* Ignore vs error
    * Please bad records for human review
* Ability to run part of ETL on bad data/process
* Privacy? (GDPR)
* Discussion
    * What are you doing?
    * What tools are you using? (mention ones from PyCon talk?)
    * Process
* great expectations Python package
