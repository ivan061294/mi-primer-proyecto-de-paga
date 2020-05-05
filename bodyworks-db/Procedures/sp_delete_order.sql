DELIMITER $$
CREATE PROCEDURE `sp_delete_order` (IN pId INT,
								  OUT coderror INT,
								  OUT msgerror VARCHAR(20))
BEGIN
	DECLARE v_status VARCHAR(20);
	SET coderror = 0;
	SET msgerror = 'OK';
	SELECT status INTO v_status FROM orders WHERE id = pId;
	IF v_status = 'E' THEN
		DELETE FROM orders WHERE id = pId;
		DELETE FROM orderdetails WHERE order_id = pId;
	ELSEIF v_status IS NULL THEN
		SET coderror = 1;
		SET msgerror = CONCAT('No existe la orden de trabajo #', pId);
	ELSE
		SET coderror = 1;
		SET msgerror = CONCAT('No se puede eliminar, el estado (', v_status, ') no es valido');
	END IF;
	SELECT coderror, msgerror;
END$$
DELIMITER ;