-- Seed apps and products table
INSERT INTO apps (id, name, icon, created_at, updated_at)
VALUES (
    1,
    'SaaS Todos',
    'https://portal.its.ac.id/images/icon-app/myITS-StudentConnect.png',
    now(),
    now()
  ),
  (
    2,
    'SaaS Notes',
    'https://portal.its.ac.id/images/icon-app/myITS-StudentConnect.png',
    now(),
    now()
  ),
  (
    3,
    'SaaS Image App',
    'https://portal.its.ac.id/images/icon-app/myITS-StudentConnect.png' now(),
    now()
  );
INSERT INTO products (
    id,
    app_id,
    tier_name,
    price,
    deployment_schema,
    created_at,
    updated_at
  )
VALUES (
    '2cac359e-d35a-458c-8474-3b7b9cc3984c',
    1,
    'Basic',
    10000,
    '{
  "deployment_repository_url": "https://github.com/godlixe/saas-todo.git",
  "terraform_execution_path": "saas-todo/terraform/tiers/1",
  "init_script_path": "saas-todo/terraform/tiers/1/my_script.go",
  "migration_script_path":"",
  "deployment_model": "pool"
}',
    now(),
    now()
  ),
  (
    '908c22d0-0b7a-4820-a500-27440ab73c63',
    1,
    'Premium',
    20000,
    '{
  "deployment_repository_url": "https://github.com/godlixe/saas-todo.git",
  "terraform_execution_path": "saas-todo/terraform/tiers/1",
  "init_script_path": "saas-todo/terraform/tiers/1/my_script.go",
  "migration_script_path":"",
  "deployment_model": "pool"
}',
    now(),
    now()
  ),
  (
    '63900f28-1d1d-4df1-a448-abc1c2756326',
    1,
    'Platinum',
    30000,
    '{
  "deployment_repository_url": "https://github.com/godlixe/saas-todo.git",
  "terraform_execution_path": "saas-todo/terraform/tiers/1",
  "init_script_path": "saas-todo/terraform/tiers/1/my_script.go",
  "migration_script_path":"",
  "deployment_model": "pool"
}',
    now(),
    now()
  ),
  (
    '81e54445-80f6-43c4-8c64-378cae02a74f',
    2,
    'First',
    1,
    '{
  "deployment_repository_url": "https://github.com/godlixe/saas-todo.git",
  "terraform_execution_path": "saas-todo/terraform/tiers/1",
  "init_script_path": "saas-todo/terraform/tiers/1/my_script.go",
  "migration_script_path":"",
  "deployment_model": "pool"
}',
    now(),
    now()
  ),
  (
    '4900214b-bd86-48b4-b3a4-94dc67007372',
    2,
    'Second',
    2,
    '{
  "deployment_repository_url": "https://github.com/godlixe/saas-todo.git",
  "terraform_execution_path": "saas-todo/terraform/tiers/1",
  "init_script_path": "saas-todo/terraform/tiers/1/my_script.go",
  "migration_script_path":"",
  "deployment_model": "pool"
}',
    now(),
    now()
  ),
  (
    '780b3ee2-60f1-430a-9c09-fdba1816a3ab',
    3,
    'Basic',
    5,
    '{
  "deployment_repository_url": "https://github.com/godlixe/saas-todo.git",
  "terraform_execution_path": "saas-todo/terraform/tiers/1",
  "init_script_path": "saas-todo/terraform/tiers/1/my_script.go",
  "migration_script_path":"",
  "deployment_model": "pool"
}',
    now(),
    now()
  ),
  (
    '164c8302-2146-4f1c-82c3-b93f1d7e41dd',
    3,
    'Pro',
    10,
    '{
  "deployment_repository_url": "https://github.com/godlixe/saas-todo.git",
  "terraform_execution_path": "saas-todo/terraform/tiers/1",
  "init_script_path": "saas-todo/terraform/tiers/1/my_script.go",
  "migration_script_path":"",
  "deployment_model": "pool"
}',
    now(),
    now()
  )