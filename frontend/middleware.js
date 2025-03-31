import { NextResponse } from "next/server";
const AUTH_PAGES = ["/login", "/"];
const isAuthPages = (url) => AUTH_PAGES.some((page) => page.startsWith(url));
export async function middleware(request) {
    const { url, nextUrl, cookies } = request;
    const { value: token } = cookies.get("social-network-jwt") ?? {
        value: null,
    };

    const isAuthPageRequested = isAuthPages(nextUrl.pathname);
    const isJWTValid = await (
        await fetch(`${"http://backend:8000"}/api/verify`, {
            headers: {
                Authorization: token,
            },
        })
    ).json();

    if (isAuthPageRequested) {
        if (isJWTValid.code !== 200) {
            const response = NextResponse.next();
            response.cookies.delete("social-network-jwt");
            return response;
        }
        const response = NextResponse.redirect(new URL(`/app/home`, url));
        return response;
    }

    if (isJWTValid.code !== 200) {
        const searchParams = new URLSearchParams(nextUrl.searchParams);
        searchParams.set("next", nextUrl.pathname);
        const response = NextResponse.redirect(
            new URL(`/login?${searchParams}`, url)
        );
        response.cookies.delete("token");
        return response;
    }
    return NextResponse.next();
}

export const config = { matcher: ["/login", "/", "/app/:path*"] };
