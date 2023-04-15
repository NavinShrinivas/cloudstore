import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';

function ProductCard({ product, AddtoOrder }) {
    return <Card style={{ width: '20rem', margin: '1rem' }} className="shadow-sm">
        <Card.Body>
            <Card.Title>{product.name}</Card.Title>
            <Card.Text>
                {product.manufacturer}
            </Card.Text>
            <Card.Text>
                {product.avg_rating ? (
                    <>
                        {
                            Array.from({ length: Math.ceil(product.avg_rating) }, (_, i) =>
                                <i key={i} className="bi bi-star-fill" />)
                        }
                        {
                            Array.from({ length: 5 - Math.ceil(product.avg_rating) }, (_, i) =>
                                <i key={i} className="bi bi-star" />)
                        }
                    </>
                ) : null}
            </Card.Text>
            <Card.Text>
                {product.price} $ / {product.limit} Items left
            </Card.Text>
            <Card.Text>
                Sold by {product.seller_username}
            </Card.Text>
            <Button value={product.name} onClick={AddtoOrder} variant="primary">Order</Button>
        </Card.Body>
    </Card>;
}

export default ProductCard;
