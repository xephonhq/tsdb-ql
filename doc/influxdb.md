# InfluxDB

InfluxDB is a greate reference for tsdb-ql

## TODO

- [ ] infux shell does not support tab complete https://github.com/influxdata/influxdb/issues/7704
- [ ] infux ql said you can't distinguish field and tag

## Install

- Download binary from https://www.influxdata.com/downloads/#influxdb
- Extract and put the bin folder in path, i.e. `export PATH=$PATH:$HOME/app/influxdb/usr/bin`
- run `influxd` to start the daemon
- run `influx` to enter interactive shell or import data
- [ ] TODO: figure out where the disk data is stored

## Query Language

The entry can be found in https://docs.influxdata.com/influxdb/v1.2/query_language/

- `curl https://s3.amazonaws.com/noaa.water-database/NOAA_data.txt -o NOAA_data.txt`
- `SELECT COUNT(water_level) FROM h2o_feet`
- `SELECT * FROM h2o_feet LIMIT 5`
- [InfluxQl Short](https://docs.influxdata.com/influxdb/v1.2/query_language/data_exploration/#the-basic-select-statement)

### Select

- `SELECT <field_key>[,<field_key>,<tag_key>] FROM <measurement_name>[,<measurement_name>]`

### Where

- field value can be string/value
- tag value are all string
- timestamp
  - [ ] TODO: we can have a strict mode?
  - [x] does it support now? now+1min etc. Yes it does, `now() + 6m`

- [ ] functions

### Group by

- https://vimeo.com/200898048
- does not work with fields
- `SELECT "water_level" FROM "h2o_feet" GROUP BY "location"`
