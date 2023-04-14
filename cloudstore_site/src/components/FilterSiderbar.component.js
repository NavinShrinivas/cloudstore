import { useState, useEffect } from 'react';

function FilterSidebar(props) {
    props = {
        filter: {
            price: [0, Infinity],
            ratings: [0, Infinity],
            seller: []
        }
    }

    const priceRanges = [[0, 100], [100, 200], [200, 300], [300, 400], [400, 500]]
    const [currentPriceRange, setCurrentPriceRange] = useState([0, Infinity])
    const setPrice = (event) => {
        // setPriceRange(event.target.value)
        const priceRange = event.target.value.split(',')
        setCurrentPriceRange([priceRange[0], priceRange[1]])
        console.log(currentPriceRange[0])
        console.log(currentPriceRange[1])

    }



    return (
        <div className="filter-sidebar m-2 p-2 border ">

            <h1>Filter</h1>
            <div>
                <h2>Price</h2>
                <div>
                    <input name="price" type="radio" value={[0, Infinity]} onChange={setPrice} defaultChecked />
                    <label> All</label>
                </div>
                {priceRanges.map((range) => {
                    return (
                        <div key={range[0]}>
                            <input name="price" type="radio" value={range} onChange={setPrice} />
                            <label> {range[0]} - {range[1]}</label>
                        </div>
                    )
                })}
                <div>
                    <input name="price" type="radio" value={[priceRanges[priceRanges.length - 1][1], Infinity]} onChange={setPrice} />
                    <label> {priceRanges[priceRanges.length - 1][1]}+</label>
                </div>
                <div>
                    <input name="price" type="number" value={currentPriceRange[0]} onChange={setPrice} style={{ width: '100px' }} />
                    &nbsp;&nbsp;
                    <input name="price" type="number" value={currentPriceRange[1]} onChange={setPrice} style={{ width: '100px' }} />
                </div>
            </div>
            <br />
            <div>
                <h2>Ratings</h2>

            </div>
            <br />
            <div>
                <h2>Seller</h2>
            </div>
        </div>
    )

}

export default FilterSidebar;