drop table if exists t;
create table t(a int, b varchar);
insert into t values(1, null);
insert into t values(2, "abc");
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
insert into t select * from t;
select count(*) from t;
-- @sleep:10
select count(*) from metadata_scan('table_func_metadata_scan.t', '*') g;
select count(*) from metadata_scan('table_func_metadata_scan.t', 'a') g;
select column_name, rows_count, null_count, compress_size, origin_size from metadata_scan('table_func_metadata_scan.t', 'a') g;
select column_name, rows_count, null_count, compress_size, origin_size from metadata_scan('table_func_metadata_scan.t', '*') g;
select sum(compress_size) from metadata_scan('table_func_metadata_scan.t', '*') g;
select sum(origin_size) from metadata_scan('table_func_metadata_scan.t', '*') g;
drop table if exists t;