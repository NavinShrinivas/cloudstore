import axios from 'axios';
import { useSelector } from 'react-redux'
import { useState, useEffect } from 'react';

import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';

import { HomeNavbar, FilterSidebar } from '../components'

function BuyerHomePage() {


    const [user, setUser] = useState(useSelector((state) => state.user.value))
    const loginurl = "http://localhost:5001"
    const produrl = "http://localhost:5002"

    const [data, setData] = useState([])
    const [order, setOrder] = useState([])

    axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 

    useEffect(() => {
        getproducts()
    }, [])


    const AddtoOrder = (event) => {
        order.push(event.target.value)
        // li_order.push(event.target.value)
        console.log(order)
    }

    const getproducts = () => {
        axios.get(produrl + "/api/products/all", {}).then(function (resp) {
            setData(resp.data.products)
        })
    }
    return (
        <>
            <div className="home-page" style={{ backgroundColor: '#f5f5f5', height: '100vh', overflow: 'hidden' }}>
                <HomeNavbar />
                <div className="mx-4" style={{ marginTop: '70px' }}>

                    <div className="row">
                        <div className="col-3">
                            <FilterSidebar />
                        </div>
                        <div className="col-9">
                            <div className="row">
                                {
                                    data.map(user => {
                                        return (
                                            <Card style={{ minWidth: '8rem', margin: '1rem' }}>
                                                <Card.Body>
                                                    <Card.Title>{user.name}</Card.Title>
                                                    <Card.Text>
                                                        {user.age}
                                                    </Card.Text>
                                                    <Card.Text>
                                                        {user.manufacturer}
                                                    </Card.Text>
                                                    <Card.Text>
                                                        {user.price} $ / {user.limit} Items left
                                                    </Card.Text>
                                                    <Button key={user.name} value={user.name} onClick={AddtoOrder} variant="primary" >Order</Button>
                                                </Card.Body>
                                            </Card>
                                        )
                                    })}
                            </div>
                        </div>
                        <div>{order}</div>
                    </div>
                </div>
            </div >
        </>
    );
}

export default BuyerHomePage;
