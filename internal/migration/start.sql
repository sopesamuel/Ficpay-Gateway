CREATE TABLE IF NOT EXISTS payments (
    payment_id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(64) NOT NULL,
    customer_id VARCHAR(64) NOT NULL,
    amount INT NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status ENUM('PENDING','AUTHORIZED','CAPTURED','VOIDED','REFUNDED') NOT NULL,
    auth_id VARCHAR(36) NULL,
    capture_id VARCHAR(36) NULL,
    void_id VARCHAR(36) NULL,
    refund_id VARCHAR(36) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)

-- update auto refreshes after there is a change, the "on update" i mean.

CREATE TABLE IF NOT EXISTS state_history (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    payment_id VARCHAR(36) NOT NULL, 
    from_status ENUM('PENDING','AUTHORIZED','CAPTURED','VOIDED','REFUNDED') NULL, 
    to_status ENUM('PENDING','AUTHORIZED','CAPTURED','VOIDED','REFUNDED') NOT NULL, 
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_id) REFERENCES payments(payment_id)
)