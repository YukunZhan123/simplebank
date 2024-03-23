ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS accounts_owner_currency_idx;
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS accounts_owner_fkey;
DROP TABLE users;