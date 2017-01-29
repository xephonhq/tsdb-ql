# Specification draft

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
