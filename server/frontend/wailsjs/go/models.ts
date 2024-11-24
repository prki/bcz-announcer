export namespace main {
	
	export class CalloutUpdateMessage {
	    callout_id: string;
	    update_request: string;
	
	    static createFrom(source: any = {}) {
	        return new CalloutUpdateMessage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.callout_id = source["callout_id"];
	        this.update_request = source["update_request"];
	    }
	}
	export class GGOBRow {
	    player_name: string;
	    points: number;
	
	    static createFrom(source: any = {}) {
	        return new GGOBRow(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.player_name = source["player_name"];
	        this.points = source["points"];
	    }
	}
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
	export class Player {
	    player_name: string;
	    country_code: string;
	
	    static createFrom(source: any = {}) {
	        return new Player(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.player_name = source["player_name"];
	        this.country_code = source["country_code"];
	    }
	}
	export class SetupMatch {
	    game_name: string;
	    p1_name: string;
	    p1_country_code: string;
	    p2_name: string;
	    p2_country_code: string;
	    callout_id: string;
	
	    static createFrom(source: any = {}) {
	        return new SetupMatch(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.game_name = source["game_name"];
	        this.p1_name = source["p1_name"];
	        this.p1_country_code = source["p1_country_code"];
	        this.p2_name = source["p2_name"];
	        this.p2_country_code = source["p2_country_code"];
	        this.callout_id = source["callout_id"];
	    }
	}

}

