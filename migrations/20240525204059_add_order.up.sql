CREATE TABLE orders
(
    id uuid NOT NULL,
    customer_id uuid,
    items JSON[] DEFAULT ARRAY[]::JSON[],
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT pk_order_id PRIMARY KEY (id),
    CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customer (id),
    CONSTRAINT status CHECK (status IN ('PENDING', 'PAID', 'REJECTED', 'AWAITING_PAYMENT'))
);
