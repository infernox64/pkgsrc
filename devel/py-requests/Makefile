# $NetBSD: Makefile,v 1.39 2018/06/16 13:08:32 adam Exp $

DISTNAME=	requests-2.19.1
PKGNAME=	${PYPKGPREFIX}-${DISTNAME}
CATEGORIES=	devel www python
MASTER_SITES=	${MASTER_SITE_PYPI:=r/requests/}

MAINTAINER=	imil@NetBSD.org
HOMEPAGE=	http://www.python-requests.org/
COMMENT=	HTTP library, written in Python, for human beings
LICENSE=	apache-2.0

DEPENDS+=	${PYPKGPREFIX}-OpenSSL>=0.14:../../security/py-OpenSSL
DEPENDS+=	${PYPKGPREFIX}-certifi>=2017.4.17:../../security/py-certifi
DEPENDS+=	${PYPKGPREFIX}-chardet>=3.0.2:../../converters/py-chardet
DEPENDS+=	${PYPKGPREFIX}-cryptography>=1.3.4:../../security/py-cryptography
DEPENDS+=	${PYPKGPREFIX}-idna>=2.5:../../www/py-idna
DEPENDS+=	${PYPKGPREFIX}-urllib3>=1.12.1:../../www/py-urllib3
TEST_DEPENDS+=	${PYPKGPREFIX}-Socks>=1.5.8:../../net/py-Socks
TEST_DEPENDS+=	${PYPKGPREFIX}-test>=2.8.0:../../devel/py-test
TEST_DEPENDS+=	${PYPKGPREFIX}-test-cov-[0-9]*:../../devel/py-test-cov
TEST_DEPENDS+=	${PYPKGPREFIX}-test-httpbin-[0-9]*:../../www/py-test-httpbin
TEST_DEPENDS+=	${PYPKGPREFIX}-test-mock-[0-9]*:../../devel/py-test-mock
TEST_DEPENDS+=	${PYPKGPREFIX}-test-xdist-[0-9]*:../../devel/py-test-xdist

USE_LANGUAGES=		# none
REPLACE_PYTHON+=	requests/certs.py

.include "../../lang/python/application.mk"
.include "../../lang/python/egg.mk"
.include "../../mk/bsd.pkg.mk"
