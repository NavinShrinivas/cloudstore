import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
import axios from 'axios';
import { useNavigate } from "react-router-dom";

function ProductSellerCard({ product }) {
   const navigate = useNavigate();
   const DeleteProduct = (e) => {
      axios.delete("/api/products/delete", {
         data: {
            id: product.ID
         }
      }).then(function(resp) {

         console.log(resp)
         if (resp.data.status) {
            alert("Product deleted!")
            navigate("/")
            window.location.reload()
         } else {
            alert("Error deleting product :(")
            navigate("/")
         }
      }).catch(function(err) {
         console.log(err)
         alert("Error deleting product :(")
         navigate("/")
      });

   }
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
         <Button value={product.ID} onClick={DeleteProduct} variant="warn">Delete</Button>
      </Card.Body>
   </Card>;
}

export default ProductSellerCard;
