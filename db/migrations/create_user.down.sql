DROP TRIGGER IF EXISTS trg_users_set_updated_at ON users;
DROP FUNCTION IF EXISTS set_updated_at;
DROP TABLE IF EXISTS users;
DO $$
BEGIN
  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
DROP TYPE user_status;
END IF;
  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
DROP TYPE user_role;
END IF;
END$$;
