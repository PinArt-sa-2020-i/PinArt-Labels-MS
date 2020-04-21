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
  `idLabel` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `created_on` DATETIME NOT NULL DEFAULT NOW(),
  `description` VARCHAR(45) NULL,
  PRIMARY KEY (`idLabel`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `labels`.`Label_relation`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Label_relation` ;

CREATE TABLE IF NOT EXISTS `labels`.`Label_relation` (
  `idLabel_relation` INT NOT NULL AUTO_INCREMENT,
  `Label_id1` INT NOT NULL,
  `Label_idLabel` INT NOT NULL,
  PRIMARY KEY (`idLabel_relation`, `Label_id1`, `Label_idLabel`),
  CONSTRAINT `fk_Label_relation_Label`
    FOREIGN KEY (`Label_id1`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Label_relation_Label1`
    FOREIGN KEY (`Label_idLabel`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_Label_relation_Label_idx` ON `labels`.`Label_relation` (`Label_id1` ASC) VISIBLE;

CREATE INDEX `fk_Label_relation_Label1_idx` ON `labels`.`Label_relation` (`Label_idLabel` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `labels`.`Board_Label`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Board_Label` ;

CREATE TABLE IF NOT EXISTS `labels`.`Board_Label` (
  `idBoard_Label` INT NOT NULL AUTO_INCREMENT,
  `created_on` DATETIME NOT NULL DEFAULT NOW(),
  `Label_id` INT NOT NULL,
  `Board_id` INT NOT NULL,
  PRIMARY KEY (`idBoard_Label`, `Label_id`, `Board_id`),
  CONSTRAINT `fk_Board_Label_Label1`
    FOREIGN KEY (`Label_id`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Board_Label_Board1`
    FOREIGN KEY (`Board_id`)
    REFERENCES `labels`.`Board` (`idBoard`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_Board_Label_Label1_idx` ON `labels`.`Board_Label` (`Label_id` ASC) VISIBLE;

CREATE INDEX `fk_Board_Label_Board1_idx` ON `labels`.`Board_Label` (`Board_id` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `labels`.`Label_User`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `labels`.`Label_User` ;

CREATE TABLE IF NOT EXISTS `labels`.`Label_User` (
  `idLabel_User` INT NOT NULL AUTO_INCREMENT,
  `created_on` DATETIME NOT NULL DEFAULT NOW(),
  `User_idUser` INT NOT NULL,
  `Label_idLabel` INT NOT NULL,
  PRIMARY KEY (`idLabel_User`, `User_idUser`, `Label_idLabel`),
  CONSTRAINT `fk_Label_User_User1`
    FOREIGN KEY (`User_idUser`)
    REFERENCES `labels`.`User` (`idUser`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Label_User_Label1`
    FOREIGN KEY (`Label_idLabel`)
    REFERENCES `labels`.`Label` (`idLabel`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_Label_User_User1_idx` ON `labels`.`Label_User` (`User_idUser` ASC) VISIBLE;

CREATE INDEX `fk_Label_User_Label1_idx` ON `labels`.`Label_User` (`Label_idLabel` ASC) VISIBLE;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- -----------------------------------------------------
-- Data for table `labels`.`Label`
-- -----------------------------------------------------
START TRANSACTION;
USE `labels`;
INSERT INTO `labels`.`Label` (`idLabel`, `name`, `created_on`, `description`) VALUES (, 'carros', DEFAULT, 'los carros ');
INSERT INTO `labels`.`Label` (`idLabel`, `name`, `created_on`, `description`) VALUES (DEFAULT, 'Paisajes', DEFAULT, 'vistas fotograficas');
INSERT INTO `labels`.`Label` (`idLabel`, `name`, `created_on`, `description`) VALUES (DEFAULT, 'Perros', DEFAULT, 'Perros');
INSERT INTO `labels`.`Label` (`idLabel`, `name`, `created_on`, `description`) VALUES (DEFAULT, 'Gatos', DEFAULT, 'Gatos');
INSERT INTO `labels`.`Label` (`idLabel`, `name`, `created_on`, `description`) VALUES (DEFAULT, 'Tatuajes M', DEFAULT, 'Tatuajes masculinos');
INSERT INTO `labels`.`Label` (`idLabel`, `name`, `created_on`, `description`) VALUES (DEFAULT, 'Herramientas', DEFAULT, 'Herramientas de contruccion');
INSERT INTO `labels`.`Label` (`idLabel`, `name`, `created_on`, `description`) VALUES (DEFAULT, 'Computadores', DEFAULT, 'Equipos gamer');

COMMIT;

