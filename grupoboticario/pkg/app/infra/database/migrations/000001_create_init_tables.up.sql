-- -----------------------------------------------------
-- Table public.RETAILER
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS public.RETAILER (
  ID_RETAILER SERIAL PRIMARY KEY,
  NAME_RETAILER VARCHAR(160) NOT NULL,
  CPF VARCHAR(200) NOT NULL UNIQUE,
  EMAIL VARCHAR(200) NOT NULL,
  PSW VARCHAR(300) NOT NULL
);
CREATE INDEX CPF_INDEX ON public.RETAILER (CPF);
-- -----------------------------------------------------
-- Table public.CREDIT
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS public.CREDIT (
  ID_CREDIT SERIAL PRIMARY KEY,
  COMMISION FLOAT NOT NULL,
  IS_ACTIVE BOOLEAN NOT NULL
);
-- -----------------------------------------------------
-- Table public.PURCHASE
-- -----------------------------------------------------
CREATE TYPE status_purchase AS ENUM ('On Validation', 'Approved');
CREATE TABLE IF NOT EXISTS public.PURCHASE (
  CODE VARCHAR(45) PRIMARY KEY,
  AMMOUNT DECIMAL(10, 2) NOT NULL,
  DATA DATE NOT NULL DEFAULT CURRENT_DATE,
  STATUS_PURCHASE status_purchase,
  FK_RETAILER_ID_RETAILER INT NOT NULL,
  FK_CREDIT_ID_CREDIT INT NOT NULL,
  CONSTRAINT fk_PURCHASE_RETAILER
    FOREIGN KEY (FK_RETAILER_ID_RETAILER)
    REFERENCES public.RETAILER (ID_RETAILER)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT fk_CREDIT_ID_CREDIT
  FOREIGN KEY (FK_CREDIT_ID_CREDIT)
  REFERENCES public.CREDIT (ID_CREDIT)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION
);
-- -----------------------------------------------------
-- INSERT CREDIT
-- -----------------------------------------------------
INSERT INTO public.CREDIT (COMMISION, IS_ACTIVE)
VALUES (0.10, TRUE),
  (0.15, TRUE),
  (0.20, TRUE);