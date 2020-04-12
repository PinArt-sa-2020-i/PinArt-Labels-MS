-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema labels
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `labels` ;

-- -----------------------------------------------------
-- Schema labels
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `labels` ;
USE `labels` ;

-- -----------------------------------------------------
-- Table `labels`.`User`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`User` ;

CREATE TABLE IF NOT EXISTS `labels`.`User` (
  `idUser` INT NOT NULL,
  PRIMARY KEY (`idUser`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `labels`.`Board`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Board` ;

CREATE TABLE IF NOT EXISTS `labels`.`Board` (
  `idBoard` INT NOT NULL,
  PRIMARY KEY (`idBoard`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `labels`.`Label`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Label` ;

CREATE TABLE IF NOT EXISTS `labels`.`Label` (
  `idLabel` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `createdAt` DATE NULL,
  `description` VARCHAR(45) NULL,
  PRIMARY KEY (`idLabel`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `labels`.`Label_relation`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Label_relation` ;

CREATE TABLE IF NOT EXISTS `labels`.`Label_relation` (
  `idLabel_relation` INT NOT NULL,
  `Label_id1` INT NOT NULL,
  `Label_idLabel` INT NOT NULL,
  PRIMARY KEY (`idLabel_relation`, `Label_id1`, `Label_idLabel`),
  INDEX `fk_Label_relation_Label_idx` (`Label_id1` ASC) VISIBLE,
  INDEX `fk_Label_relation_Label1_idx` (`Label_idLabel` ASC) VISIBLE,
  CONSTRAINT `fk_Label_relation_Label`
    FOREIGN KEY (`Label_id1`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Label_relation_Label1`
    FOREIGN KEY (`Label_idLabel`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `labels`.`Board_Label`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Board_Label` ;

CREATE TABLE IF NOT EXISTS `labels`.`Board_Label` (
  `idBoard_Label` INT NOT NULL,
  `createdAt` DATE NULL,
  `Label_id` INT NOT NULL,
  `Board_id` INT NOT NULL,
  PRIMARY KEY (`idBoard_Label`, `Label_id`, `Board_id`),
  INDEX `fk_Board_Label_Label1_idx` (`Label_id` ASC) VISIBLE,
  INDEX `fk_Board_Label_Board1_idx` (`Board_id` ASC) VISIBLE,
  CONSTRAINT `fk_Board_Label_Label1`
    FOREIGN KEY (`Label_id`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Board_Label_Board1`
    FOREIGN KEY (`Board_id`)
    REFERENCES `labels`.`Board` (`idBoard`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `labels`.`Label_User`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Label_User` ;

CREATE TABLE IF NOT EXISTS `labels`.`Label_User` (
  `idLabel_User` INT NOT NULL,
  `createdAt` DATE NULL,
  `User_idUser` INT NOT NULL,
  `Label_idLabel` INT NOT NULL,
  PRIMARY KEY (`idLabel_User`, `User_idUser`, `Label_idLabel`),
  INDEX `fk_Label_User_User1_idx` (`User_idUser` ASC) VISIBLE,
  INDEX `fk_Label_User_Label1_idx` (`Label_idLabel` ASC) VISIBLE,
  CONSTRAINT `fk_Label_User_User1`
    FOREIGN KEY (`User_idUser`)
    REFERENCES `labels`.`User` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Label_User_Label1`
    FOREIGN KEY (`Label_idLabel`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
