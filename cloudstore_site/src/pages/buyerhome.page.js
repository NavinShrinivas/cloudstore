import axios from 'axios';
// import { useSelector } from 'react-redux'
import { useState, useEffect } from 'react';

import { HomeNavbar, ProductCard, FilterSidebar } from '../components'

function BuyerHomePage() {
    // const [user, setUser] = useState(useSelector((state) => state.user.value))

    const [filter, setFilter] = useState({
        search: '',
        price: [0, Infinity],
        ratings: 0,
        seller: [],
        manufacturer: []
    })

    const [products, setProducts] = useState([])
    const [productsData, setProductsData] = useState([])

    // eslint-disable-next-line
    const [order, setorder] = useState([])
    // eslint-disable-next-line
    const [sellers, setsellers] = useState(new Set())
    // eslint-disable-next-line
    const [manufacturers, setmanufacturers] = useState(new Set())

    useEffect(() => {
        getproducts()
        // eslint-disable-next-line
    }, [])

    useEffect(() => {
        setProducts(productsData.filter((product) => {
            return product.name.toLowerCase().includes(filter.search.toLowerCase()) &&
                parseInt(product.price) >= parseInt(filter.price[0]) &&
                (filter.price[1] === Infinity || parseInt(filter.price[1]) === 0 ? true : parseInt(product.price) <= parseInt(filter.price[1])) &&
                parseInt(product.avg_rating) >= parseInt(filter.ratings) &&
                (filter.seller.length === 0 || filter.seller.includes(product.seller_username)) &&
                (filter.manufacturer.length === 0 || filter.manufacturer.includes(product.manufacturer))
        }))
    }, [filter, productsData])

    const AddtoOrder = (event) => {
        order.push(event.target.value)
        // li_order.push(event.target.value)
        console.log(order)
    }

    axios.defaults.withCredentials = true //NOTE : This is very important to be able to set cookies 

    const getproducts = () => {
        axios.get("/api/products/all", {}).then(function (resp) {
            setProductsData(resp.data.products)
            setProducts(resp.data.products)
            for (let i = 0; i < resp.data.products.length; i++) {
                sellers.add(resp.data.products[i].seller_username)
                manufacturers.add(resp.data.products[i].manufacturer)
            }
            setFilter({
                ...filter,
                seller: [...sellers],
                manufacturer: [...manufacturers]
            })
        })
    }

    return (
        <>
            <div className="home-page" style={{ backgroundColor: '#f5f5f5', overflow: 'hidden' }}>
                <HomeNavbar />
                <div className="mx-4" style={{ marginTop: '4rem' }}>
                    <div className="row">
                        <div className="col-3">
                            <FilterSidebar sellers={sellers} manufacturers={manufacturers} filter={filter} setFilter={setFilter} />
                        </div>
                        <div className="col-9">
                            <div className="row">
                                {
                                    products.map(product => {
                                        return (
                                            <ProductCard key={product.id} product={product} AddtoOrder={AddtoOrder} />
                                        )
                                    })
                                }
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
