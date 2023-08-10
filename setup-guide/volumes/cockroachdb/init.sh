#!/bin/sh
cockroach sql --url  postgresql://my_node_user:#!N@d$P@sswo@rd@0.0.0.0:26257/my_node_db?sslmode=verify-full
 --file ./node/dbinit.sql