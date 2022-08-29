select
  -- Required Columns
  arn as resource,
  case
    when package_type <> 'Zip' then 'skip'
    when runtime in ('nodejs16.x',  'python3.9',  'ruby2.7',  'go1.x', 'java11') then 'ok'
    else 'alarm'
  end as status,
  case
    when package_type <> 'Zip' then title || ' package type is ' || package_type || '.'
    when runtime in ('nodejs16.x',  'python3.9',  'ruby2.7',  'go1.x', 'java11') then title || ' uses latest runtime - ' || runtime || '.'
    else title || ' uses ' || runtime || ' which is not the latest version.'
  end as reason,
  -- Additional Dimensions
  region,
  account_id
from
  aws_lambda_function;
