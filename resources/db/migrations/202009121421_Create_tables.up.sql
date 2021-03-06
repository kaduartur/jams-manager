CREATE TABLE CATEGORIES (
    CATEGORY_ID VARCHAR(50) NOT NULL PRIMARY KEY,
    NAME VARCHAR(100) NOT NULL,
    CREATED_AT TIMESTAMP WITH TIME ZONE NOT NULL,
    UPDATED_AT TIMESTAMP WITH TIME ZONE
);

CREATE TABLE RIDERS (
  RIDER_ID VARCHAR(50) NOT NULL PRIMARY KEY,
  NAME VARCHAR(100) NOT NULL,
  AGE VARCHAR(10) NOT NULL,
  GENDER VARCHAR(100) NOT NULL,
  CITY VARCHAR(100) NOT NULL,
  CPF VARCHAR(20) NOT NULL,
  PAID_SUBSCRIPTION boolean,
  SPONSORS VARCHAR(200) NOT NULL,
  CATEGORY_ID VARCHAR(50) REFERENCES CATEGORIES(CATEGORY_ID),
  CREATED TIMESTAMP WITH TIME ZONE NOT NULL,
  UPDATED TIMESTAMP WITH TIME ZONE
);

CREATE TABLE SCORES (
    SCORE_ID VARCHAR(50) NOT NULL PRIMARY KEY,
    RIDER_ID VARCHAR(50) REFERENCES RIDERS(RIDER_ID),
    SCORE DECIMAL NOT NULL,
    CREATED_AT TIMESTAMP WITH TIME ZONE NOT NULL,
    UPDATED_AT TIMESTAMP WITH TIME ZONE
)