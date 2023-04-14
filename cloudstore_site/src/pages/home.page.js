import { useState } from 'react'
import { useSelector } from 'react-redux'

import { BuyerHomePage, SellerHomePage } from './'

function HomePage() {

    // eslint-disable-next-line
    const [user, setUser] = useState(useSelector((state) => state.user.value))

    return (
        <>
            {(user && (user.usertype) ?
                <>
                    {(user.usertype === 'buyer') ? <BuyerHomePage /> : null}
                    {(user.usertype === 'seller') ? <SellerHomePage /> : null}
                </>
                : null)}
        </>
    );
}

export default HomePage;
