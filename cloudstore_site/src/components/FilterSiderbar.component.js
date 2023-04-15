import { useState } from 'react';

function FilterSidebar(props) {

    const [customPriceRange, setCustomPriceRange] = useState(['', ''])
    const priceRanges = [[0, 100], [100, 200], [200, 300], [300, 400], [400, 500]]

    const onChangePriceInput = (e) => {
        const re = /^[0-9\b]+$/;
        if (e.target.value === '' || e.target.value === null) {
            setCustomPriceRange(['', ''])
        }
        if (re.test(e.target.value)) {
            e.target.value = parseInt(e.target.value)
            if (e.target.id === 'pricemin') {
                if (e.target.value >= customPriceRange[1]) {
                    setCustomPriceRange([e.target.value, e.target.value])
                } else {
                    setCustomPriceRange([e.target.value, customPriceRange[1]])
                }
            } else {
                if (e.target.value <= customPriceRange[0]) {
                    setCustomPriceRange([e.target.value, e.target.value])
                } else {
                    setCustomPriceRange([customPriceRange[0], e.target.value])
                }
            }
        }
    }

    const applyPrice = () => {
        if (customPriceRange[0] === '') {
            props.setFilter({
                ...props.filter,
                price: [0, customPriceRange[1]]
            })
        } else if (customPriceRange[1] === '') {
            props.setFilter({
                ...props.filter,
                price: [customPriceRange[0], Infinity]
            })
        } else {
            props.setFilter({
                ...props.filter,
                price: [customPriceRange[0], customPriceRange[1]]
            })
        }
    }

    const setSearch = (event) => {
        props.setFilter({
            ...props.filter,
            search: event.target.value
        })
    }

    const setPrice = (event) => {
        setCustomPriceRange(['', ''])
        const priceRange = event.target.value.split(',')
        if (priceRange[1] === 'Infinity') {
            priceRange[1] = Infinity
        }
        props.setFilter({
            ...props.filter,
            price: [priceRange[0], priceRange[1]]
        })
    }

    const setRatings = (event) => {
        props.setFilter({
            ...props.filter,
            ratings: event.target.value
        })
    }

    const setSeller = (event) => {
        if (event.target.checked) {
            props.setFilter({
                ...props.filter,
                seller: [...props.filter.seller, event.target.value]
            })
        } else {
            props.setFilter({
                ...props.filter,
                seller: props.filter.seller.filter((seller) => seller !== event.target.value)
            })
        }
    }

    const setManufacturer = (event) => {
        if (event.target.checked) {
            props.setFilter({
                ...props.filter,
                manufacturer: [...props.filter.manufacturer, event.target.value]
            })
        } else {
            props.setFilter({
                ...props.filter,
                manufacturer: props.filter.manufacturer.filter((manufacturer) => manufacturer !== event.target.value)
            })
        }
    }

    return (
        <div className="position-fixed filter-sidebar m-2 p-3 card" style={{ height: '85vh', maxHeight: '85vh', width: '23%', overflowY: 'scroll' }}>

            <input className="form-control" placeholder='Search' onChange={setSearch} />
            <br />
            <div>
                <h4>Price</h4>
                <div>
                    <input name="price" type="radio" value={[0, Infinity]} onChange={setPrice} defaultChecked />
                    <label>
                        &nbsp;
                        No filter
                    </label>
                </div>
                {priceRanges.map((range) => {
                    return (
                        <div key={range[0]}>
                            <input name="price" type="radio" value={range} onChange={setPrice} />
                            &nbsp;

                            <label> {range[0]} - {range[1]}</label>
                        </div>
                    )
                })}
                <div>
                    <input name="price" type="radio" value={[priceRanges[priceRanges.length - 1][1], Infinity]} onChange={setPrice} />
                    <label>
                        &nbsp;
                        {priceRanges[priceRanges.length - 1][1]}+
                    </label>
                </div>
                &nbsp;
                <div className="d-flex">
                    <input name="price" id="pricemin" className="form-control" placeholder='MIN' value={customPriceRange[0]} style={{ width: '30%' }} onChange={onChangePriceInput} />
                    &nbsp;&nbsp;
                    <input name="price" id="pricemax" className="form-control" placeholder='MAX' value={customPriceRange[1]} style={{ width: '30%' }} onChange={onChangePriceInput} />
                    &nbsp;&nbsp;
                    <button className="btn btn-outline-secondary" onClick={applyPrice}>Apply</button>
                </div>
            </div>
            <br />
            <div>
                <h4>Ratings</h4>
                <input name="rating" type='radio' value={0} defaultChecked onChange={setRatings} />
                <label>
                    &nbsp;
                    No filter
                </label>


                {[4, 3, 2, 1].map((stars) => {
                    return (
                        <div key={stars}>
                            <input name="rating" type='radio' value={stars} onChange={setRatings} />
                            <label>
                                &nbsp;

                                {Array.from({ length: stars }, (_, i) =>
                                    <i key={i} className="bi bi-star-fill" />)}
                                {Array.from({ length: 5 - stars }, (_, i) =>
                                    <i key={i} className="bi bi-star" />)}
                                &nbsp;
                                & up
                                &nbsp;
                            </label>
                        </div>
                    )
                })}
            </div>
            <br />
            {
                props.sellers ?
                    <div>
                        <h4>Seller</h4>
                        {props.sellers && Array.from(props.sellers).map((seller) => {
                            return (
                                <div key={seller}>
                                    <input name="seller" type='checkbox' defaultChecked value={seller} onChange={setSeller} />
                                    <label>
                                        &nbsp;
                                        {seller}
                                    </label>
                                </div>
                            )
                        })}
                    </div>
                    : null
            }
            <br />
            {
                props.manufacturers ?
                    <div>
                        <h4>Manufacturer</h4>
                        {props.manufacturers && Array.from(props.manufacturers).map((manufacturer) => {
                            return (
                                <div key={manufacturer}>
                                    <input name="manufacturer" type='checkbox' defaultChecked value={manufacturer} onChange={setManufacturer} />
                                    <label>
                                        &nbsp;
                                        {manufacturer}
                                    </label>
                                </div>
                            )
                        })}
                    </div>
                    : null
            }


        </div >
    )

}

export default FilterSidebar;