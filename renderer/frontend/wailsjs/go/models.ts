export namespace main {
	
	export class Message {
	    action: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Message(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.action = source["action"];
	        this.message = source["message"];
	    }
	}

}

