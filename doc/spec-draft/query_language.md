# Query Language Draft

## Initial version

This is what is presented in the proposal, which has data generation and query ability

````
INSERT INTO "docker_monitor" (metric, tags, value, timestamp)
  VALUES ("cpu.usage",
          {
            "provider": rand in ["aws", "azure", "aliyun", "gce"],
            "region": rand in ["us", "eu", "cn"],
            "os": rand in ["fedora", "win"]
          }),
          rand in [1, 2...100]
          [now, now+1s, ...])
--- synthetic data for cpu usage of docker images across multiple dc and providers          
````

````
SELECT "cpu.usage" FROM "docker_monitor"
  WHERE "provider" = "aliyun" AND "region" = "cn" AND "os" = "win"
    AND timestamp BETWEEN now AND now+1d
--- grab cpu usage for all the win images running on aliyun cn region
````

## Naive version

Split into three parts

- fake data generation
- query (like normal SQL)
  - query optimization
- graph

The first and third part are kind of like R I guess

## Ref

### SparkQL

- grammar file https://github.com/apache/spark/blob/master/sql/catalyst/src/main/antlr4/org/apache/spark/sql/catalyst/parser/SqlBase.g4
- https://databricks.com/blog/2015/04/13/deep-dive-into-spark-sqls-catalyst-optimizer.html
