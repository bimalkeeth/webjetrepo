import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from "@angular/common/http";
import {Observable} from "rxjs/internal/Observable";
import {Injectable} from "@angular/core";
@Injectable()
export class TokenInterceptor implements HttpInterceptor {

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    req = req.clone({ headers: req.headers.append('x-access-token', 'sjd1HfkjU83ksdsm3802k') });
    req = req.clone({ headers: req.headers.append('Access-Control-Allow-Origin', 'localhost:5000') });

    return next.handle(req);
  }
}
